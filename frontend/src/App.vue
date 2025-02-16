<script lang="ts" setup>
import { computed, onMounted, reactive } from 'vue';
import { useExecutor } from './http/executor';
import { todoApi } from './http/api';
import { IListItem, IState } from './types.dto';

const state = reactive<IState>({
  list: [],
  newTask: ''
})

const getList = computed(() => state.list.sort((a, b) => +a.is_done - +b.is_done))

async function addTask() {
  if (!state.newTask) {
    alert("Значение не должно быть пустым")
    return 
  }
  const newTodo = {title: state.newTask}

  await useExecutor().run({
    request: todoApi.add(newTodo),
    onResult: (json: any) => {
      console.log(json, "json")
      alert(json['result'])
    },
    onComplete: () => {
      fetchListTodo()
      state.newTask = ''
    }
  })
}

async function toggleTask(index: number) {
  const currentTask = state.list[index]
  currentTask.is_done = !currentTask.is_done
  const {id, ...restTask} = currentTask
  await useExecutor().run({
    request: todoApi.update(id, restTask),
    onResult: (json: any) => {
      console.log(json, "json")
    },
    onComplete: () => {
      fetchListTodo()
    }
  })
}

async function removeTask(id: number) {
  const isConfirmed = window.confirm("Вы уверены, что хотите удалить этот элемент?");
  if (!isConfirmed) return 
  
  await useExecutor().run({
    request: todoApi.delete(id),
    onResult: (json: any) => {
      console.log(json, "json")
    },
    onComplete: () => {
      fetchListTodo()
    }
  })
}

onMounted(() => fetchListTodo())

async function fetchListTodo() {
  await useExecutor().run({
    request: todoApi.list(),
    onResult: (json: any) => {
      state.list = json
    }
  })
}
</script>

<template>
  <div id="app" class="container">
    <h2>Список задач</h2>
    <div class="add_task">
      <UiInput 
        v-model="state.newTask" 
        @keyup.enter="addTask" 
        placeholder="Введите задачу" 
      />
      <UiButton  @click="addTask">Добавить</UiButton>
    </div>
    <ul>
      <li v-for="(task, index) in getList" :key="`${task.title}-${index}`" class="task" :class="{ done: task.is_done }">
        <span>{{ task.id }}){{ task.title }}</span>
        <div class="btns">
          <UiButton type="secondary" @click="toggleTask(index)">{{task.is_done ? "Убрать" : "Завершено"}}</UiButton>
          <UiButton  type="deleted" @click="removeTask(task.id)">Удалить</UiButton>
        </div>
      </li>
    </ul>
  </div>
</template>

<style>
body {
  font-family: Arial, sans-serif;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f4f4f4;
}
.container {
  height: 80%;
  width: 400px;
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}
.task {
  color: black;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  border-bottom: 1px solid #ddd;
}
.task.done {
  text-decoration: line-through;
  color: gray;
}
.btns {
  display: flex;
  gap: 3px;
}
.add_task {
  width: 80%;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 3fr 1fr;
  gap: 5px;
  justify-content: center;
}
</style>
