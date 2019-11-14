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
  font-family: Arial;

  button {
    background: orange;
    border: none;
    padding: 1rem;
    color: white;
    font-weight: bold;
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
