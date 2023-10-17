# What is Autovivification?

[Autovivification](https://en.wikipedia.org/wiki/Autovivification) can be found in many interpreted languages, and some compiled ones.

It's the creation of the intermediary layers in nested structures.

The article in Wikipedia has some good examples, here's one for Perl:

```perl
%h = ();          # %h is an empty hash
$h{A}{B}{C}{D}=1  # creates %h = (A => {B => {C => {D => 1}}})
```

# What is this package?

`avmap` adds autovivification to the maps only. It arose from a task where several nested maps were needed
to keep intermediary results of a calculation.

Given the declaration

```go
type calcs struct {
    someAttrs  map[string]int
    otherAttrs map[string]int
    ...
}
perStateMap := map[string]*calcs{}
```

It allows you to switch from this

```go
for ... {
    state := ...
    someAttr := ...
    otherAttr := ...
    if _, ok := perStateMap[state]; !ok {
        perStateMap[state] = &calcs{
            someAttrs:  map[string]int{},
            otherAttrs: map[string]int{},
            ...
        }
    }
    if _, ok := perStateMap[state].someAttrs[someAttr]; !ok {
        perStateMap[state].someAttrs[someAttr] = 0
    }
    perStateMap[state].someAttrs[someAttr]++
    if _, ok := perStateMap[state].otherAttrs[otherAttr]; !ok {
        perStateMap[state].otherAttrs[otherAttr] = 0
    }
    perStateMap[state].otherAttrs[otherAttr]++
}
```

to this

```go
for name := range ... {
    state := ...
    someAttr := ...
    otherAttr := ...
    avmap.SetIfMissing(perStateMap, state, &calcs{
            someAttrs:  map[string]int{},
            otherAttrs: map[string]int{},
            ...
        }
    )
    avmap.Inc(perStateMap[state].someAttrs, someAttr)
    avmap.Inc(perStateMap[state].otherAttrs, otherAttr)
}
```
