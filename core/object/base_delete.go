package object

import (
	"os"
	"path/filepath"
	"strings"

	"opensvc.com/opensvc/core/objectactionprops"
	"opensvc.com/opensvc/util/file"
)

type OptsDelete struct {
	OptsGlobal
	OptsLock
	RID         string `flag:"rid"`
	Unprovision bool   `flag:"unprovision"`
}

//
// Delete is the 'delete' object action entrypoint.
//
// If no resource selector is set, remove all etc, var and log
// file belonging to the object.
//
// If a resource selector is set, only delete the corresponding
// sections in the configuration file.
//
func (t Base) Delete(options OptsDelete) error {
	props := objectactionprops.Delete
	unlock, err := t.lockAction(props, options.OptsLock)
	if err != nil {
		return err
	}
	defer unlock()
	return t.lockedDelete(options)
}

func (t Base) lockedDelete(opts OptsDelete) error {
	if opts.RID != "" {
		return t.deleteSections(opts.RID)
	}
	return t.deleteInstance()
}

func (t Base) deleteInstance() error {
	if err := t.deleteInstanceFiles(); err != nil {
		return err
	}
	if err := t.deleteInstanceLogs(); err != nil {
		return err
	}
	if err := t.setPurgeCollectorTag(); err != nil {
		t.log.Warn().Err(err).Msg("")
		return nil
	}
	return nil
}

func (t Base) deleteInstanceFiles() error {
	patterns := []string{
		filepath.Join(t.logDir(), t.Path.Name+".log*"),
		filepath.Join(t.logDir(), t.Path.Name+".debug.log*"),
		filepath.Join(t.logDir(), "."+t.Path.Name+".log*"),
		filepath.Join(t.logDir(), "."+t.Path.Name+".debug.log*"),
		filepath.Join(t.varDir()),
	}
	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			t.log.Warn().Err(err).Str("pattern", pattern).Msg("expand glob for delete")
			continue
		}
		for _, fpath := range matches {
			_ = t.tryDeleteInstanceFile(fpath)
		}
	}
	t.tryDeleteInstanceFile(t.ConfigFile())
	return nil
}

func (t Base) tryDeleteInstanceFile(fpath string) bool {
	if file.IsProtected(fpath) {
		t.log.Warn().Str("path", fpath).Msg("block attempt to remove a protected file")
		return false
	}
	if err := os.RemoveAll(fpath); err != nil {
		t.log.Warn().Err(err).Str("path", fpath).Msg("removing")
		return false
	}
	t.log.Info().Str("path", fpath).Msg("removed")
	return true
}

func (t Base) deleteInstanceLogs() error {
	return nil
}

func (t Base) setPurgeCollectorTag() error {
	return nil
}

func (t Base) deleteSections(s string) error {
	sections := strings.Split(s, ",")
	return t.config.DeleteSections(sections)
}
