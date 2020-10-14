
This program uses selected algorithms from the book "Astronomical Algorithms" by Jean Meeus, as implemented in go packages by Sonia keys

https://github.com/soniakeys/meeus

Thhe above can be downloaded and installed on your system with
<pre>
go get -t github.com/soniakeys/meeus/...
</pre>

The environment variable VSOP87 must be set to point at the directory containing the VSOP87B* data files e.g.
<pre>
export VSOP87=/Users/chris/src/go-area/astro/vsop87
</pre>

Usage
<pre>
Usage of conj.go
  -end_year int
    	Start Year (default 2020)
  -planet1 string
    	First Planet (default "Jupiter")
  -planet2 string
    	Second Planet (default "Saturn")
  -start_year int
    	Start Year (default 2020)
</pre>

Example run
<pre>
$ go run conj.go -start_year 2000 -planet1 Saturn -planet2 Jupiter -end_year 2020 
 
-----------------------------------------------------------------
Detecting conjunctions between Saturn and Jupiter 

Start Date: 1 January 2000
End Date: 31 December 2020
 
O  Planets greater than 3 degrees apart. Time step is: 1 days.
o  Planets less than 3 degrees apart. Time step is: 1.992 hours.
.  Planets less than 0.3 degrees apart. Time step is: 2.88 minutes.

Only print event details if planets less than 0.3 degrees apart
-----------------------------------------------------------------
O o O o . 
-----------------------------------------------------------------
Saturn <-> Jupiter
2020 December 21 18:24 UTC
Minimum separation: 6′6″
-----------------------------------------------------------------
o 

Done!
</pre>
