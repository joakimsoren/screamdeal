<template>
  <div @drop.prevent class="file-upload-container">
    <div class="no-theme-selected" v-if="!themeSelected">
      <h3>Please select a theme to upload a file</h3>
    </div>
    <label
      v-else
      for="file-input"
      class="file-upload"
      @drop.prevent="handleInput($event.dataTransfer.files[0])"
      @dragover.prevent
      @dragenter.prevent
      @dragleave.prevent
    >
      <span v-if="file">
        <h2>Uploading:</h2>
        <h3>{{ file.name }}</h3>
      </span>
      <h2 v-else>Click here or drop to upload file</h2>
    </label>
    <input
      id="file-input"
      @input="handleInput($event.target.files[0])"
      type="file"
    />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component
export default class FileUpload extends Vue {
  @Prop() themeSelected: boolean;
  @Prop() loading: boolean;
  @Prop() loaded: boolean;

  file: File = undefined;

  handleInput(file: File) {
    this.file = file;
    this.$emit("file", file);
  }
}
</script>
<style scoped lang="scss">
.no-theme-selected,
.file-upload {
  height: 10rem;
  width: 30%;
  margin-left: auto;
  margin-right: auto;
  background: rgba(0, 0, 0, 0.5);
  border: 2px dashed rgb(68, 68, 68);
  display: block;
  color: darkorange;
  padding: 1rem;
  h3 {
    font-weight: 200;
  }
}
.no-theme-selected {
  border-style: solid;
}
input {
  position: absolute;
  opacity: 0;
}
</style>
