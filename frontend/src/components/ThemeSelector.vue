<template>
  <div class="theme-selector">
    <div
      v-for="item in themes"
      :key="item"
      :class="['theme-box', { selected: item === selected }]"
    >
      <img class="theme-image" src="/img/doc.jpg" />
      <button class="theme-button" @click="handleClick(item)">
        Select theme
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component
export default class ThemeSelector extends Vue {
  @Prop() selected: string;
  @Prop() themes: [string];

  handleClick(index: string) {
    this.$emit("select", index);
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
      box-shadow: 0px 5px 10px 0px #2e2e2e;
      border: 1px solid transparent;
    }

    &.selected > .theme-image {
      border: 1px solid green;
      box-shadow: 0px 5px 10px 0px green;
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
