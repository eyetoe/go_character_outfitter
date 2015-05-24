package main

import (
    "fmt"
    "math/rand"
    "time"
    "net/http"
)

func main() {

    http.HandleFunc("/", handler) // redirect all urls to the handler function
    http.ListenAndServe(":667", nil) // listen for connections at port 9999 on the local machine
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Inside handler")
    fmt.Fprintln(w, "<html>")

    fmt.Fprintln(w, "<head>")
    fmt.Fprintln(w, "<link rel='stylesheet' type='text/css' href='http://xonk.org/xonk.css' />")
    fmt.Fprintln(w, "</head>")

    fmt.Fprintln(w, "<center><table border='1' width='85%' align=center><tr><td align='left'>")

    fmt.Fprintln(w, "<h1>Random Character Outfitter</h1>")
    fmt.Fprintln(w, "<p><i>(refresh for a new hero)</i></p>")
    fmt.Fprintf(w, "<hr>")

    fmt.Fprintf(w, "<p><b>Name: </b>"+get_FirstName()+" "+get_LastName()+"</p>")
    fmt.Fprintf(w, "<p><b>Melee: </b>"+get_Melee()+"</p>")
    fmt.Fprintf(w, "<p><b>Ranged: </b>"+get_Ranged()+"</p>")
    fmt.Fprintf(w, "<p><b>Long Range: </b>"+get_Long()+"</p>")
    fmt.Fprintf(w, "<p><b>Gernade: </b>"+get_Gernade()+"</p>")
    fmt.Fprintf(w, "<hr>")
    fmt.Fprintf(w, "<p><i>It was a cold forboding morning and <b>"+get_FirstName()+"</b> was inspecting his <b>"+get_Melee()+"</b>, hoping for some action.  Just as papa <b>"+get_LastName()+"</b> had taught.  You can't be too prepared.</p><p>Suddenly <b>"+get_FirstName()+"</b> spied a <b>"+get_Monster()+"</b> coming over the horizon, but still out of range of his <b>"+get_Long()+"</b>.  It wasn't surprised and emitted a loud aggressive sound.</p><p>As the <b>"+get_Monster()+"</b> charged, <b>"+get_FirstName()+"</b> lobbed a <b>"+get_Gernade()+"</b> gernade toward the foe.  The creature was momentarily stunned while <b>"+get_FirstName()+"</b> pulled out his <b>"+get_Ranged()+"</b> and aimed at the beast.  Three shots were not enought to deter it, as the <b>"+get_Monster()+"</b> came ever closer.</p><p>Finally the beast arrived at close range.  With a quick stroke of the <b>"+get_Melee()+"</b> the <b>"+get_Monster()+"</b> took it right in the <b>"+get_Part()+"</b> and died in a lump.  This would be a good tale to tell over at the tavern tonight.  Three cheers for <b>"+get_FirstName()+"</b> <b>"+get_LastName()+"</b></p>")

    fmt.Fprintln(w, "</td></tr></table></center>")
    fmt.Fprintln(w, "</html>")
}

func get_Part() string{

    rand.Seed(time.Now().Unix())

    // slice
    weapons := []string{
                        "jejunum",
                        "esophagus",
                        "lips",
                        "eyelashes",
                        "pinky toe",
                        "shin",
                        "nipple",
                        "uvula",
                        "everything",
                        "vestigial tail",
                        "nether regions",
                        "left gonad",
                        "twig and berries",
                        "elbow",
                        "spleen",
                        "junk",
                        "butt cheek",
                        }

    // return random element
    return weapons[rand.Intn(len(weapons))]
}

func get_Monster() string{

    rand.Seed(time.Now().Unix())

    // slice
    weapons := []string{
                        "Skeletal Racoon",
                        "Humplefrump",
                        "Calimari",
                        "Unicorn on the cob",
                        "Jabberwocky",
                        "Odd Sock",
                        "Skeevy Skivvies",
                        "Bandersnatch",
                        "12ft Duck",
                        "Tinysaurus Rex",
                        "Slithering Slime",
                        "Owl Badger",
                        "Mole Rat",
                        "Earthworm",
                        "Tardigrade",
                        "Kangashrew",
                        }

    // return random element
    return weapons[rand.Intn(len(weapons))]
}

func get_Melee() string{

    rand.Seed(time.Now().Unix())

    // slice
    weapons := []string{
                        "Sword",
                        "Sledgehammer",
                        "Chainsaw",
                        "Welding Torch",
                        "Ice Pick",
                        "Crowbar",
                        "Wrench",
                        "Old Boot",
                        "Broom",
                        "Stun Stick",
                        "Fart",
                        "Pokey Stick",
                        }

    // return random element
    return weapons[rand.Intn(len(weapons))]
}

func get_Ranged() string{

    rand.Seed(time.Now().Unix())

    // slice
    weapons := []string{
                        "Flame Thrower",
                        "Shotgun",
                        "Double Barrel Shotgun",
                        "Gernade Launcher",
                        "Throwing Knife",
                        "Tazer Stun Gun",
                        "Short Bow",
                        "Crossbow Pistol",
                        }

    // return random element
    return weapons[rand.Intn(len(weapons))]
}

func get_Long() string{

    rand.Seed(time.Now().Unix())

    // slice
    weapons := []string{
                        "Pistol",
                        "SMG",
                        "Assault Rifle",
                        "Hunting Rifle",
                        "Sniper Rifle",
                        "Rocket Launcher",
                        "Crossbow",
                        "Long Bow",
                        }

    // return random element
    return weapons[rand.Intn(len(weapons))]
}

func get_Gernade() string{

    rand.Seed(time.Now().Unix())

    // slice
    weapons := []string{
                        "Explosive",
                        "Flash Bang",
                        "Tear Gas",
                        "Bouncy",
                        "Sticky",
                        "Incindiary",
                        "Cryo",
                        "Tesla",
                        "Molotov",
                        "Cluster",
                        "Shrapnel",
                        }

    // return random element
    return weapons[rand.Intn(len(weapons))]
}

func get_FirstName() string{

    rand.Seed(time.Now().Unix())

    // slice
    names := []string{
                       "Dink",
                       "Woosh",
                       "Bark",
                       "Derp",
                       "Gern",
                       "Slim",
                       "Meh",
                       "Walnut",
                       "Walrus",
                       "Dingo",
                       "Poodle",
                       "Doofus",
                       "Flip",
                       "Sherbert",
                       "Spatula",
                       "Fecalpap",
                       "Smorgasbord",
                       "Silly",
                       "Ted",
                       "Burn",
                       "Tingle",
                       "Splash",
                       "Johnny",
                       "Tuna",
                       "Nirvana",
                       "Wap",
                        }

    // return random element
    return names[rand.Intn(len(names))]
}

func get_LastName() string{

    rand.Seed(time.Now().Unix())

    // slice
    names := []string{
                       "Waperson",
                       "Dinkski",
                       "Wooshblatz",
                       "Slime",
                       "Cashew",
                       "Blandsten",
                       "Palendrome",
                       "Fooperson",
                       "Happyslap",
                       "Koala-nose",
                       "Slothface",
                       "Wilikins",
                       "Wilson",
                       "Jammypants",
                       "Googlethat",
                       "Nevermind",
                       "Tuner",
                       "Fishface",
                       "Fists-o-fury",
                       "Baba Ghanoush",
                       "Not-in-the-face",
                        }


    // return random element
    return names[rand.Intn(len(names))]
}
