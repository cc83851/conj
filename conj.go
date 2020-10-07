package main

import (
        "fmt"
        "time"
        "math"
        "flag"

        pp "github.com/soniakeys/meeus/v3/planetposition"
        "github.com/soniakeys/meeus/v3/elliptic"
        "github.com/soniakeys/meeus/v3/angle"
        "github.com/soniakeys/meeus/v3/julian"
        "github.com/soniakeys/sexagesimal"
        "github.com/soniakeys/unit"
)

func main() {

//        var start_year int = 2020
//        var end_year int = 2020
      
//        var Planet1 string = "Jupiter"
//        var Planet2 string = "Saturn"

        var Planet1 string
        flag.StringVar(&Planet1, "planet1", "Jupiter", "First Planet")
        var Planet2 string
        flag.StringVar(&Planet2, "planet2", "Saturn", "Second Planet")
        var start_year int
        flag.IntVar(&start_year, "start_year", 2020, "Start Year")
        var end_year int
        flag.IntVar(&end_year, "end_year", 2020, "Start Year")

        flag.Parse()

        var large_incr float64 = 1.0
        var medium_incr float64 = 0.083
        var small_incr float64 = 0.002

        var large_threshold unit.Angle = 3.0/57.3
        var small_threshold unit.Angle = 0.3/57.3
        var tiny_threshold unit.Angle = 0.3/57.3 // Only print events with min sep less than this
//
//      NO CHANGES BELOW HERE
//
        planets := make(map[string]int)
        planets["Mercury"] = 0
        planets["Venus"] = 1
        planets["Mars"] = 3
        planets["Jupiter"] = 4
        planets["Saturn"] = 5
        planets["Uranus"] = 6
        planets["Neptune"] = 7

        var P1 int = planets[Planet1]  
        var P2 int = planets[Planet2] 

        var last_separation unit.Angle = math.Pi

	jd_start := julian.CalendarGregorianToJD(start_year, 1, 1)
	jd_end := julian.CalendarGregorianToJD(end_year, 12, 31)

        current := jd_start
        var found bool = false
        var large bool = false
        var medium bool = false
        var small bool = false

	earth, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
	    fmt.Println(err)
	    return
    	}

	planet1, err := pp.LoadPlanet(P1)
	if err != nil {
	    fmt.Println(err)
	    return
        }

	planet2, err := pp.LoadPlanet(P2)
	if err != nil {
	    fmt.Println(err)
	    return
	}

        fmt.Println(" ")
        fmt.Println("-----------------------------------------------------------------")
        fmt.Println("Detecting conjunctions between", Planet1, "and", Planet2, "\n")
	y, m, d := julian.JDToCalendar(jd_start)
        fmt.Printf("Start Date: %v %s %v\n", d, time.Month(m), y)
        fmt.Printf("End Date: %v %s %v\n", 31, time.Month(12), end_year)
        fmt.Println(" ")
        fmt.Println("O  Planets greater than", large_threshold*57.3, "degrees apart. Time step is:", large_incr, "days.") 
        fmt.Println("o  Planets less than", large_threshold*57.3, "degrees apart. Time step is:", medium_incr*24, "hours.") 
        fmt.Println(".  Planets less than", small_threshold*57.3, "degrees apart. Time step is:", small_incr*24*60, "minutes.") 
        fmt.Println("\nOnly print event details if planets less than", tiny_threshold*57.3, "degrees apart") 
        fmt.Println("-----------------------------------------------------------------")

        for current < jd_end {
	    r1, d1 := elliptic.Position(planet1, earth, current)
	    r2, d2 := elliptic.Position(planet2, earth, current)
            rr1 := unit.Angle(r1)
            rr2 := unit.Angle(r2)
	    separation := angle.SepHav(rr1, d1, rr2, d2)

            if separation < small_threshold {
                if small == false {
                    small = true
                    fmt.Print(". ")
                }
                if separation > last_separation {
                    if !found {
                        if last_separation < tiny_threshold {
                            fmt.Print("\n")
                            fmt.Println("-----------------------------------------------------------------")
                            fmt.Println(Planet1,"<->", Planet2)
	                    y, m, d := julian.JDToCalendar(current)
                            hour := 24*(d - float64(int(d)))
                            minute := 60*(hour - float64(int(hour)))
                            fmt.Printf("%d %s %v %v:%v UTC\n", y, time.Month(m), int(d), int(hour), int(minute))
                            sexlast := fmt.Sprint(sexa.FmtAngle(last_separation))
    	                    fmt.Println("Minimum separation:",sexlast)
//                            fmt.Println(current.utc_strftime('%Y-%m-%d %H:%M:%S'), "UTC")
//                            fmt.Println("Minimum separation:",last_separation*57.3)     
                            fmt.Println("-----------------------------------------------------------------")
                            found = true
                        }
                    }
                }
                current += small_incr
                medium,large = false, false
            } else if separation < large_threshold {
                if !medium {
                    medium = true
                    fmt.Print("o ")
                }
                current += medium_incr
                small,large = false, false
            } else {
                if large == false {
                    large = true
                    fmt.Print("O ")
                }
                current += large_incr
                small, medium = false, false
                found = false
            }

            last_separation = separation
        } 
        fmt.Println("\n\nDone!")
}
