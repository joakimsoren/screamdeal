<template>
  <div class="theme-selector">
    <div v-for="item in 5" :key="item" class="theme-box">
      <img class="theme-image" src="/img/doc.jpg" />
      <button class="theme-button" @click="handleClick(index)">Select theme</button>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class ThemeSelector extends Vue {
  handelClick(index) {
    this.$emit("selectTheme", index);
  }
}
</script>
<style scoped lang="scss">
.theme-selector {
  display: flex;
  width: 80%;
  margin-left: auto;
  margin-right: auto;

  .theme-box {
    width: 20%;
    height: auto;
    margin: 2rem;
    position: relative;
    transition: width 1s;

    .theme-image {
      width: 100%;
      height: auto;
    }

    .theme-button {
      background: orange;
      border: none;
      padding: 1rem;
      color: white;
      font-weight: bold;
      opacity: 0;
      transition: opacity 0.3s;
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      cursor: pointer;
      &:focus {
        outline: none;
      }

      &:hover {
        background: darkorange;
      }
    }

    &:hover {
      width: 50%;
      .theme-button {
        transition: opacity 0.3s 1s;
        opacity: 1;
      }
    }
  }
}
</style>
