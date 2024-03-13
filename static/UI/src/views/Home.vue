<template>
  <v-container>
    <v-row>
      <v-col
        cols="6"
        offset="3"
      >
        <div>
          <p align="center">
            <img
              src="@/assets/toolbox.png"
              height="60"
              alt="Project Logo"
            >
          </p>
          <h3 align="center">
            Toolbox
          </h3>
          <p align="center">
            A collection of various tools for developers, engineers and
            programmers
          </p>
          <p align="center">
            <a href="https://github.com/tiborhercz/toolbox/releases"><img
              src="https://img.shields.io/github/downloads/tiborhercz/toolbox/total.svg"
              alt="GitHub Downloads"
            ></a>
            <a href="https://github.com/tiborhercz/toolbox/releases/latest"><img
              src="https://img.shields.io/github/release/tiborhercz/toolbox.svg"
              alt="Latest Release"
            >
            </a>
            <a href="https://github.com/tiborhercz/toolbox/actions/workflows/go-ci.yaml"><img
              src="https://img.shields.io/github/workflow/status/tiborhercz/toolbox/Build"
              alt="Build Status"
            >
            </a>
          </p>
        </div>
      </v-col>
    </v-row>
    <v-row>
      <v-col
        cols="8"
        offset="2"
      >
        <p>
          A collection of various tools like encoding, decoding, hashing, ipv4 cidr calculation, and
          more.
          The tools can be accessed from the command line or by using the web UI.
        </p>
        <p>
          A lot of tools like these are hosted online and use backend system to do the processing.
          This means you don't know what they do with your sensitive data. Another downside of
          online tools is that they are littered with ads.
        </p>
        <p>
          This toolbox runs locally on your machine and all the processing is done in your browser
          using WebAssembly or JavaScript.
        </p>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'Home',

  components: {},
}
</script>
