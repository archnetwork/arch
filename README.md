# Arch 🏹
## It's my framework and I'll Serve() how I want to

Arch (archnet to differentiate from the operating system) is a server framework. There are hundreds of them out there, but this is mine. I like it, and I'm not building it for anyone else.

The end goal is to build something that allows edge functions like Workers/Deploy/Compute@Edge to work collaboratively with more complex servers in more traditional locales like, I dunno, a VM somewhere.

The more modest starting point is to build something that works equally well as a single instance as it does a distributed service-oriented thing.

An Arch application is a single codebase compiled into a single binary that can be deployed in a multitude of ways depending on the `dependencies` it has access to. That means that if you deploy a single instance with all the dependencies it needs, it'll run happily. If you run 10 different variants with different (sub)sets of dependencies, it'll act more like a (micro)service system, but the code will remain unchanged.

Arch is super mega ultra opinionated. Developers of an Arch app create objects called `ArchWays` to describe the **way** a request should be handled. Every request is handled in 4 stages: `Analyze, Route, Cache, Handle` (hence ARCH). Every deployed Arch server node will dynamically determine if it's capable of handling a stage, and pass execution off to another node if it does not have the required dependencies. Nodes communicate with one another over a mesh.

The end goal of a collaborative edge-and-other-servers application means that Arch will be more of a protocol than a single library or implementation. It'll be able to run in JS edge-y environments (with some help from Wasm), and in other environments by compiling some code to a binary. The implementation in this repo will be a combination of Go and JS+TinyGo-Wasm. Other implementations would be possible if you wanted to implement the protocol in Rust or something else, I won't stop you.

No promises for stability or documentation correctness until a 1.0 release is cut, if this project ever gets to that point.

Any attitude in this README should not be interpreted as discouragement from participating or asking questions, only as me no longer wanting to create software that caters to a 'general audience' and instead just create something that I enjoy and think is unique. Anyone is welcome to join in via questions, PRs, issues, etc.
