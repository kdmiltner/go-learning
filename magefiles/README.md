# Magefile

[Mage](https://magefile.org/magefiles/) gives you the ability to use mage targets to streamline 
commands and can also be used in GitHub Actions. The main benefits of using mage are:
- We don't have to maintain bash scripts in a Makefile or GitHub Actions workflows.
- We can manage logic in Go and have access to all the regular Go packages.
- Lightweight and easy to use.
- Can be used with projects that aren't written in Go.
- Easy to integrate with GitHub Actions.
- Less Google-fu (for bash) and more support from the Go community.

## Install

To install mage you can follow the [instructions here](https://magefile.org/) or simply run `brew install mage`.

### Try It!

Try running a few commands after you've installed mage! To run:
- Be sure you're in the project's root directory (e.g. go-learning).
- Run any of the targets listed below from your terminal or command line interface.

## Targets

With mage, any exported function that either has no return or just an error return is consider a `target` and can be
run with the command `mage <target>`. See [this documentation](https://magefile.org/targets/) for more details.

- **HelloCFA**  - is a `Hello, World` example that prints `Hello, Chick-fil-A!`.
    - `mage helloCFA`
- **CompileApps** - will compile a list of apps based on go.mod files found.
    - `mage compileApps`
- **Test** - will run test apps.
    - `mage test {appName}` e.g. `mage test pointing-poker` will test a single app.
    - `mage test all` will test all apps.
- **RunPython** - will run the defined python script.
    - `mage runPython` will run the ./apex/interfaces/1_python/1_python_animals.py script.

## Resources

There are many helpful resources:
- [Magefile.org](https://magefile.org/) - this is the most notable as it's maintained by the creators/maintainers of mage.
- [Other resources](https://github.com/magefile/awesome-mage) - Also maintained by the mage developers is a list of resources.