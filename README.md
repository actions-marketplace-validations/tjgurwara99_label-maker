# Label Maker

I needed very simple automatic labeller for my private use and was inspired by (auto-label)[https://github.com/Renato66/auto-label].
The auto-label is very good and if you want to use something very robust, you should use that instead. I personally can't because of my specific usecase:
I like to write labelling tags in the titles of the issues like `feat:`, `bug:` etc but auto-label uses the issue body to do that. I suggested this change to the auto-label repository and the author said that he would add that in the next release but I needed it soon for my private repositories therefore I made this.

If you like to use this or want to recommend some changes, please open an issue and I'll see what I can do.


I will be adding further changes soon so that it will also be able to handle Pull Requests - doesn't seem to be that hard now that everything works as I expect it to. I will try to find a way to run these tests locally. Since its inside a docker, my current thinking process is that I will add some environment variables and test it inside that with some custom event payloads.


Steps to implement it in your repository:

1. Create `.github/workflows` directories in your repository root if you don't already have one.
2. Create a new (or edit an existing) `.yml` file in `.github/workflows` directory and write the following in the file:
```yml
name: Labeling issue 
on:
  issues:
    types: ['opened']
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: tjgurwara99/label-maker@v0.1-Beta
        with: 
          token: ${{ secrets.GITHUB_TOKEN }}
```
3. Commit these changes and add a new issue whose title contains one of your defined label names to test if it works.
