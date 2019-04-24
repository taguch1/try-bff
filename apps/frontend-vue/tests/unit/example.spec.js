import { shallowMount } from "@vue/test-utils";
import todo from "@/components/todo.vue";

describe("todo.vue", () => {
  it("renders props.todos when passed", () => {
    // const msg = "new message";
    const todos = [
      { id: "id1", title: "title A" },
      { id: "id2", title: "title B" }
    ];
    const wrapper = shallowMount(todo, {
      propsData: { todos }
    });
    expect(wrapper.findAll("li").length).toEqual(2);
  });
});
