<template>
  <div class="theme-selector">
    <div
      v-for="(item, index) in 5"
      :key="item"
      :class="['theme-box', { selected: index === selected }]"
    >
      <img class="theme-image" src="/img/doc.jpg" />
      <button class="theme-button" @click="handleClick(index)">
        Select theme
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component
export default class ThemeSelector extends Vue {
  @Prop() selected: number;

  handleClick(index: number) {
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

  .theme-box {
    width: 20%;
    height: auto;
    margin: 2rem;
    position: relative;
    transition: width 1s;

    &.selected {
      border-top: solid 5px green;
    }

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
      bottom: 2rem;
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
