<template>
  <div id="app">
    <span>{{info}}</span>
    <todo v-bind:todos="todos" v-on:click-add="onAddClick" v-on:click-delete="onDeleteClick"></todo>
  </div>
</template>

<script>
import todo from "./components/todo.vue";
import * as store from "./store";

export default {
  components: {
    todo
  },
  name: "app",
  data() {
    return {
      info: "helloworld",
      todos: []
    };
  },
  created() {
    store.list().then(({ todos }) => {
      this.todos = todos;
    });
  },
  methods: {
    onAddClick: function({ title }) {
      store.add({ title }).then(todo => {
        this.todos.splice(this.todos.length, 0, todo);
      });
    },
    onDeleteClick: function(id) {
      store.remove(id).then(() => {
        this.todos.splice(this.todos.findIndex(todo => todo.id == id), 1);
      });
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
