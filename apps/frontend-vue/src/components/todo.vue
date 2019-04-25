<template>
  <div id="todo">
    <ul>
      <li v-for="todo in todos" v-bind:key="todo.id" v-on:click="onDelete(todo)">{{ todo.title }}</li>
    </ul>
    <input
      v-on:keyup.enter="onSubmit({title})"
      v-model="title"
      id="new-tilte"
      placeholder="new title"
    >
  </div>
</template>

<script>
export default {
  props: ["todos"],
  name: "todo",
  data: function() {
    return { title: "" };
  },
  methods: {
    onSubmit: function({ title }) {
      if (title) {
        this.$emit("submit", { title });
        this.title = "";
      }
    },
    onDelete: function({ id, title }) {
      this.$emit("delete", id);
      this.title = title;
    }
  }
};
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  /* display: inline-block; */
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
