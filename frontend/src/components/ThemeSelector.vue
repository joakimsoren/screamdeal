<template>
  <div class="theme-selector">
    <div
      v-for="item in themes"
      :key="item"
      :class="['theme-box', { selected: item === selected }]"
    >
      <div class="theme-image-box">
        <img class="theme-image" src="/img/doc.jpg" />
        <img class="theme" :src="item" />
        <button class="theme-button" @click="handleClick(item)">
          Select theme
        </button>
      </div>
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
  flex-direction: row;

  .theme-box {
    width: 20%;
    margin: 0.5rem;
    transition: width 1s;

    .theme-image-box {
      position: relative;
      flex-grow: 0;
      .theme-image {
        width: 100%;
        box-shadow: 0px 5px 10px 0px #2e2e2e;
        border: 1px solid transparent;
      }

      .theme {
        position: absolute;
        max-width: 50%;
        left: calc(2% + 10px);
        bottom: calc(2% + 10px);
      }
    }

    &.selected > .theme-image-box > .theme-image {
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
      width: 30%;
      height: max-content;
      .theme-button {
        transition: opacity 0.3s 1s;
        opacity: 1;
      }
    }
  }
}
</style>
