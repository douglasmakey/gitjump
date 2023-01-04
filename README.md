# GitJump

GitJump is a lightweight tool that makes it easy to navigate through the commits of a repository using a simple syntax. It's a very straightforward implementation that has dramatically improved my study method and made it easier to learn new topics."

## Why?

Did you know that you can use git as a powerful tool for learning? By browsing the commit history of a repository, you can gain insight into a project's development process and evolution. This method can be handy when learning about a new topic or trying to understand how a team tackled challenges along the way. Give it a try and see how exploring commit history can enhance your learning experience!

In addition to providing a detailed history of a project's development, exploring commit history can also help you learn about best practices and coding styles. By seeing how other developers structure their code and commit messages, you can pick up valuable tips and techniques you can apply to your work. And if you're working on a team, analyzing commit history can help you understand how your colleagues approach problems and communicate their solutions. So next time you want to learn something new, check out the commit history of relevant repositories - you might be surprised at what you can discover!

## Installation

To install GitJump, clone this repository and run the following command:

```bash
$ go install
```

## Usage

To use `gitjump`, simply run the gitjump command with the --goto flag and the desired number of commits to move. For example:

```bash
$ gitjump --goto 0
# Moves to the first commit in the repository's commit history

$ gitjump --goto 5
# Moves to the fifth commit in the repository's commit history

$ gitjump --goto -10
# Moves 10 commits back from the latest commit in the repository's commit history

```