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
  height: 35rem;

  .theme-box {
    width: 20%;
    height: max-content;
    margin: 2rem;
    position: relative;
    transition: width 1s;

    .theme-image {
      width: 100%;
      height: auto;
    }

    .theme-button {
      position: absolute;
      bottom: 2rem;
      left: 50%;
      transform: translate(-50%, -50%);
      opacity: 0;
      transition: opacity 0.3s;
    }

    &:hover {
      width: 50%;
      height: max-content;
      .theme-button {
        transition: opacity 0.3s 1s;
        opacity: 1;
      }
    }
  }
}
</style>
