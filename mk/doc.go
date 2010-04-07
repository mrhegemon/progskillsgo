/*

Authors: William Broza, Tym Lipari

dag reads in a mkfile, makes a directed acyclic graph and displays the order of
that graph. 

mk uses dag to read in the mkfile, the directed acyclic graph created is then 
used to build a list of commands that are fed into the the shell to be executed.

Usage:
  6/dag [-f mkfile] [command]

dag can take in a makefile or a list of file names

	6/mk [-f mkfile] [command]

mk can take in a makefile or a list of file names

*/
package documentation
