go-pathshorten
===

A go port of `pathshorten()` defined on Vim script

# Usage

```console
$ pathshorten ~/.vim/scripts/foo/bar/baz.vim
/h/b/.v/s/f/b/baz.vim
$ pathshorten '~/.vim/scripts/foo/bar/baz.vim'
~/.v/s/f/b/baz.vim
$ pathshorten .
.
$ pathshorten test.csv
test.csv
$ pathshorten foo/bar/baz
f/b/baz
```

Same as `vim --cmd 'echo pathshorten("~/.vim/scripts/foo/bar/baz.vim")' --cmd quit`

For more information:

```vim
pathshorten({expr})					*pathshorten()*
		Shorten directory names in the path {expr} and return the
		result.  The tail, the file name, is kept as-is.  The other
		components in the path are reduced to single letters.  Leading
		'~' and '.' characters are kept.  Example: >
			:echo pathshorten('~/.vim/autoload/myfile.vim')
			~/.v/a/myfile.vim ~
		It doesn't matter if the path exists or not.
```

# Installation

```console
$ go get github.com/b4b4r07/go-pathshorten
```

# Licence

MIT @ b4b4r07
