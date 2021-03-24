# Label Maker

I needed very simple automatic labeller for my private use and was inspired by (auto-label)[https://github.com/Renato66/auto-label].
The auto-label is very good and if you want to use something very robust, you should use that instead. I personally can't because of my specific usecase:
I like to write labelling tags in the titles of the issues like `feat:`, `bug:` etc but auto-label uses the issue body to do that. I suggested this change to the auto-label repository and the author said that he would add that in the next release but I needed it soon for my private repositories therefore I made this.

If you like to use this or want to recommend some changes, please open an issue and I'll see what I can do.


I will be adding further changes soon so that it will also be able to handle Pull Requests - doesn't seem to be that hard now that everything works as I expect it to. I will try to find a way to run these tests locally. Since its inside a docker, my current thinking process is that I will add some environment variables and test it inside that with some custom event payloads.