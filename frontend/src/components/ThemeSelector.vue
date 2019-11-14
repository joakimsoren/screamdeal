<template>
  <div class="theme-selector">
    <div v-for="item in 5" :key="item" class="theme-box">
      <image class="theme-image" src="/img/doc.jpg" />
      <button class="theme-button" @click="handleClick">Select theme</button>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class ThemeSelector extends Vue {
  handelClick(theme) {
    this.$emit("selectTheme", theme.id);
  }
}
</script>
<style scoped lang="scss">
.theme-selector {
  display: flex;
  width: 100%;
  .theme-box {
    width: 20%;
    height: 10rem;
    margin: 2rem;
    position: relative;
    background: white;

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
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
    }

    &:hover {
      .theme-button {
        opacity: 1;
      }
    }
  }
}
</style>
