<template>
  <div id="app">
    <router-view :key="$route.fullPath"></router-view>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class App extends Vue {}
</script>

<style lang="scss">
#app {
  text-align: center;
  height: 100%;
  width: 100%;
  padding: 0;
  margin: 0;
  font-family: "Lakki Reddy", cursive;

  button,
  .button {
    background: orange;
    border: none;
    font-size: 1.2rem;
    padding: 0.6rem 0.6rem 0.2rem 0.6rem;
    color: white;
    font-family: "Lakki Reddy", cursive;
    text-decoration: none;

    cursor: pointer;
    &:focus {
      outline: none;
    }

    &:hover {
      background: darkorange;
    }
  }
}
</style>
