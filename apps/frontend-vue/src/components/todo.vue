<template>
  <div id="todo">
    <ul>
      <li
        v-for="todo in todos"
        v-bind:key="todo.id"
        v-on:click="onDeleteClick(todo.id)"
      >{{ todo.title }}</li>
    </ul>
    <input v-model="newTitle" id="new-tilte" placeholder="new title">
    <button v-on:click="onAddClick({title:newTitle})">Add</button>
  </div>
</template>

<script>
export default {
  props: {
    todos: Array
  },
  data: function() {
    return { newTitle: "" };
  },
  name: "todo",
  methods: {
    onAddClick: function({ title }) {
      this.$emit("click-add", { title });
    },
    onDeleteClick: function(id) {
      this.$emit("click-delete", id);
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
