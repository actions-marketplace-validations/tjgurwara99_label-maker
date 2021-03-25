# Label Maker

I needed very simple automatic labeller for my private use. I like to write labelling tags in the titles of the issues like `feat:`, `bug:` etc, so this makes it easy for me to sort them in GitHub's UI.

Feel free to use this. If you would like to recommend some changes, please open an issue and I'll see what I can do.

With v0.1.2 we now have support for Pull Request labelling too. Now you can just add the desired label within the title of your Pull Request and this will automatically add the label that matched the title.

Steps to implement it in your repository:

1. Create `.github/workflows` directories in your repository root if you don't already have one.
2. Create a new (or edit an existing) `.yml` file in `.github/workflows` directory and write the following in the file:
```yml
name: Labeling issue 
on:
  issues:
    types: ['opened', 'edited']
  pull_request:
    types: ['opened', 'edited']
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: tjgurwara99/label-maker@v0.1.2
        with: 
          token: ${{ secrets.GITHUB_TOKEN }}
```
3. Commit these changes and add a new issue whose title contains one of your defined label names to test if it works.


This project was inspired by [auto-label](https://github.com/Renato66/auto-label) - a project made by [Renato99](https://github.com/Renato66) - which is a more complete (and robust) solution to automatic labelling than this project. However, my usecase was a bit different than the one currently supported by [auto-label](https://github.com/Renato66/auto-label), therefore I made this.
