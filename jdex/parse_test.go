package jdex_test

import (
	"os"
	"path"
	"testing"

	"github.com/itisrazza/rzjd/jdex"
	"github.com/stretchr/testify/assert"
)

func openFile(t *testing.T, name string) (file *os.File) {
	file, err := os.Open(path.Join("test", name))
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}

	return
}

func TestParseNoComment(t *testing.T) {
	file := openFile(t, "parse-no-comments.txt")
	index, err := jdex.ReadJdex(file)
	if err != nil {
		t.Fatalf("failed to parse file: %v", err)
		return
	}

	assert.Equal(t, index, buildJdex(
		[]jdex.Area{
			buildArea("10-19", "Life Admin", []jdex.Category{
				buildCategory("11", "Me & Other Living Things", []jdex.Entry{
					buildEntry("11.10", "Personal Records", map[string]string{
						"Location": "Proton Drive",
					}),
					buildEntry("11.11", "Birth Certs & Proof of Name", map[string]string{
						"Location": "Proton Drive",
					}),
					buildEntry("11.20", "Physical Health & Wellbeing", map[string]string{}),
					buildEntry("11.21", "Health Insurance & Claims", map[string]string{}),
				}),
				buildCategory("12", "Household", []jdex.Entry{
					buildEntry("12.10", "Home Records", map[string]string{}),
					buildEntry("12.11", "Official Documents", map[string]string{}),
					buildEntry("12.12", "Home Insurance", map[string]string{}),
					buildEntry("12.12+SNX", "Home Insurance (Southern Cross)", map[string]string{}),
					buildEntry("12.13", "Moving", map[string]string{
						"Location": "Google Sheets",
					}),
				}),
			}),
		},
	))
}
