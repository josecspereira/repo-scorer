# Git Scorer
This go project will, for a given CSV file, classify multiple repositories by activity,
outputting the top 10 repositories.

To execute it simply run `make run` while assuring a `commits.csv` exists in the root folder.

## The Classification System
The criteria used for the algorithm will be as follows:
* Recent commits carry more weight over older commits;
* Lines removed are preferred over lines added;
* Repositories with more unique active users have more value (and we ignore unknown users);
* More files changed are beneficial.

### Formulas
* Age factor (per commit): $e^{{Decay}*{TimeDiff}}$
    * Exponential with a decaying factor, this benefits recent commits over old ones with a very small decay;
    * Note: The decay is set to a very low number since the timestamps are roughly 3 years old
* File factor (per commit): $1.0+log(1+{FilesChanged})$
    * Diminishing returns, we want to reward multiple file changes
      but not so much that a big refactor through the whole project would jump the score massively
* Line factor (per commit): $1.0+1.0*log(1+{LinesAdded})+1.2*log(1+{LinesRemoved})$
    * Similar to file factor, except we want to slightly benefit removing lines over deleting lines
* Contributor factor (per repo): $1.0+log(1.0+{UniqueContributors})$
    * Similar to file and lines, but only account for unique contributors

### Choices & Limitations
* Currently, we are reading all the CSV file information and storing it, this is potentially a risky approach but was
  done for sake of simplicity;
* We value both additions and deletions, but an alternative approach could be using the delta of additions and deletions;
* We deliberately ignore users without a name on our active user formula, but we do consider the commit for everything else;
* Commit age was a bit tricky due to how old the commits in the CSV are, we would prefer even more recent commits;
* Testing is a bit frail as it stands, changing the formula will require fairly convoluted math
  to correctly set expectation (specially on a TDD approach);
* A lot of rounding is happening due to every commit calculation being added,
  this also means there's a fair amount of errors that propagate.