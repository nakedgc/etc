# List of nonfatal errors

(locations are approximate)

unused_vars:

	./src/cmd/compile/internal/gc/y.go:1203:
		_ = yypt // guard against "declared and not used"
		(note: does go itself hack around the 'declared and not used' error?)
	./src/cmd/compile/internal/gc/walk.go:55:
		Yyerror("%v declared and not used", l.N.Sym)
	./src/cmd/compile/internal/gc/walk.go:61:
		Yyerror("%v declared and not used", l.N.Sym)

unused_packages:

	./src/cmd/compile/internal/gc/subr.go:346:
		yyerrorl(int(pack.Lineno), "imported and not used: %q", opkg.Path)
	./src/cmd/compile/internal/gc/lex.go:2583:
		yyerrorl(int(lineno), "imported and not used: %q", path)
	./src/cmd/compile/internal/gc/lex.go:2585:
		yyerrorl(int(lineno), "imported and not used: %q as %s", path, name)


# Other notes

I wrote a `YynonFatalError` function and a `yynonFatalErrorl` function, both in `src/cmd/compile/internal/gc/subr.go`. These should be used to replace the usual `Yyerror` and `yyerrorl` calls for nonfatal errors.

The command option for the 'go' executable `-Wno-error` is handled in `src/cmd/go/build.go:224`. It is then passed onto the compiler in `src/cmd/go/build.go:392`.

The compiler catches the option `-nonfatalwarnings=1` in `src/cmd/compile/internal/gc/lex.go:228` and sets the variable `nonfatalwarnings`. The variable is then used in the nonFatalError functions mentioned above.

We also count the number of warnings for no reason whatsoever currently, using the variable `nwarnings` initialized in `src/cmd/compile/internal/gc/go.go:477`.
