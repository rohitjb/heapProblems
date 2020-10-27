package main

suspicious_env_keys = [
    "passwd",
    "password",
    "secret",
    "key",
    "access",
    "api_key",
    "apikey",
    "token",
]

pkg_update_commands = [
    "apk upgrade",
    "apt-get upgrade",
    "dist-upgrade",
]

image_tag_list = [
    "latest",
    "LATEST",
]

# Looking for suspicious environment variables
deny[msg] {
    input[i].Cmd == "env"
    val := input[i].Value
    contains(lower(val[_]), suspicious_env_keys[_])
    msg = sprintf("Suspicious ENV key found: %s \n https://sysdig.com/blog/7-docker-security-vulnerabilities/", [val])
}

# Looking for latest docker image used
warn[msg] {
    input[i].Cmd == "from"
    val := split(input[i].Value[0], ":")
    count(val) == 1
    msg = sprintf("Do not use latest tag with image: %s \n kindly refer https://vsupalov.com/docker-latest-tag/ and https://medium.com/@mccode/the-misunderstood-docker-tag-latest-af3babfd6375 for more info", [val])
}

# Looking for latest docker image used
warn[msg] {
    input[i].Cmd == "from"
    val := split(input[i].Value[0], ":")
    contains(val[1], image_tag_list[_])
    msg = sprintf("Do not use latest tag with image: %s \n kindly refer https://vsupalov.com/docker-latest-tag/ and https://medium.com/@mccode/the-misunderstood-docker-tag-latest-af3babfd6375 for more info", [input[i].Value])
}

# Looking for apk upgrade command used in Dockerfile
deny[msg] {
    input[i].Cmd == "run"
    val := concat(" ", input[i].Value)
    contains(val, pkg_update_commands[_])
    msg = sprintf("Do not use upgrade commands: %s \n kindly refer https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#run for info", [val])
}

# Looking for ADD command instead using COPY command
deny[msg] {
    input[i].Cmd == "add"
    val := concat(" ", input[i].Value)
    msg = sprintf("Use COPY instead of ADD: %s \n kindly refer https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#add-or-copy for more info", [val])
}

# sudo usage
deny[msg] {
    input[i].Cmd == "run"
    val := concat(" ", input[i].Value)
    contains(lower(val), "sudo")
    msg = sprintf("Avoid using 'sudo' command: %s \n kindly refer https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#user for more info", [val])
}