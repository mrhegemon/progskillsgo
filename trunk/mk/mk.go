/*
mk implements all of the mk.go interfaces to build the directed acyclic 
graph that governs the command order of a makefile.

Authors: William Broza, Tym Lipari
*/
package dag

import 	("os"
				"io/ioutil"
				"bytes"
				"strings"
				//"strconv"
)
