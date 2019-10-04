notabar(5) [""["notabar-VERSION"]]

# NAME

notabar - because bars are for noobs

# DESCRIPTION

The configuration files for notabar are plain text cvs files where each line is
an entry with a key and one or more values. The configuration files are read
from STDIN by notabar(1). You may create several different configuration files
and run them seperately. Below are all the possible keys and values they may
have.

Keys
	- txt,string
	- cmd,command,argument(s)

Blank lines or lines starting with \# are ignored. Newlines are removed from
commands outputs. For the txt key there are a few "special" characters allowed
which are listed below.

txt special characters
	- \\a = U+0007 = alert or bell
	- \\b = U+0008 = backspace
	- \\f = U+000C = form feed
	- \\n = U+000A = line feed or newline
	- \\r = U+000D = carriage return
	- \\t = U+0009 = horizontal tab
	- \\v = U+000b = vertical tab
	- \\\\ = U+005c = backslash
	- \\' = U+0027 = single quote
	- \\" = U+0022 = double quote

That list comes directly from golang so a current up to date list can be found
here https://golang.org/ref/spec#Rune_literals

# EXAMPLES

\#Time notification++
txt,time:++
cmd,date,+%H:%M++
txt,\n

This will print "time:20:34" if it is currently 8:34pm.

\#Volume notification++
txt,volume:++
cmd,pamixer,--get-volume-human++
txt,\n

This will print the "volume:50" if your current volume as seen 
by `pamixer --get-human-volume` is 50.

# SEE ALSO

notabar(1)

# AUTHORS

Maintained by Dakota Walsh <kota at nilsu.org>.
Up-to-date sources can be found at https://git.sr.ht/~kota/notabar/

# COPYRIGHT

Copyright 2019 Dakota Walsh

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation
and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors
may be used to endorse or promote products derived from this software without
specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
