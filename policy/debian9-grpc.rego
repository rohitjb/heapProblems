package main

disallowed_tags := ["latest"]
disallowed_req := ["https"]

denylist = [
  "add",
]

deny[msg] {
        input[i].Cmd == "from"
        val := input[i].Value
        tag := split(val[i], ":")[1]
        contains(tag, disallowed_tags[_])

        msg = sprintf("[%s] tag is not allowed", [tag])
}

deny[msg] {
        input[i].Cmd == "arg"
        val := input[i].Value

        tag := split(val[i], "=")[1]
        contains(tag, disallowed_tags[_])
        msg = sprintf("%s tag is not allowed in args", val)
}

deny[msg] {
   input[i].Cmd == "add"
   val := concat(" ", input[i].Value)

   msg = sprintf("Use COPY instead of ADD: %s", [val])
}
