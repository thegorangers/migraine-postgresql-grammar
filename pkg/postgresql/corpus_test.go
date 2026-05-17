package parser_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/thegorangers/migraine-postgresql-grammar/pkg/postgresql"
)

func TestCorpus_StableProductions_ParseClean(t *testing.T) {
	err := filepath.WalkDir("../../test/corpus/stable", func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(p, ".sql") {
			return err
		}
		t.Run(filepath.Base(p), func(t *testing.T) {
			data, err := os.ReadFile(p)
			if err != nil {
				t.Fatal(err)
			}
			errs := tryParse(string(data))
			if len(errs) > 0 {
				t.Errorf("stable fixture %s failed parse:\n%v", p, errs)
			}
		})
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCorpus_UnsupportedProductions_ProduceErrors(t *testing.T) {
	// Inverse: these MUST fail to parse cleanly. If one starts parsing
	// successfully, either the grammar grew support (move to stable/)
	// or there's a silent acceptance bug (parse error swallowed).
	err := filepath.WalkDir("../../test/corpus/unsupported", func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(p, ".sql") {
			return err
		}
		t.Run(filepath.Base(p), func(t *testing.T) {
			data, _ := os.ReadFile(p)
			errs := tryParse(string(data))
			if len(errs) == 0 {
				t.Errorf("unsupported fixture %s parsed cleanly — promote to stable/ if intentional", p)
			}
		})
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCorpus_Incomplete_KnownGaps(t *testing.T) {
	t.Skip("incomplete fixtures are tracked but not asserted; see docs/stability.md and future docs/coverage.md")
}

// tryParse runs lexer+parser, collects all syntax errors via a listener.
type errCollector struct {
	*antlr.DefaultErrorListener
	errs []string
}

func (e *errCollector) SyntaxError(_ antlr.Recognizer, _ interface{}, line, col int, msg string, _ antlr.RecognitionException) {
	e.errs = append(e.errs, msg)
}

func tryParse(sql string) []string {
	is := antlr.NewInputStream(sql)
	lexer := parser.NewPostgreSQLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPostgreSQLParser(stream)

	coll := &errCollector{}
	p.RemoveErrorListeners()
	p.AddErrorListener(coll)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(coll)

	_ = p.Root()
	return coll.errs
}
