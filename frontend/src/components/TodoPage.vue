<template>
    <div class="container">
        <h2>Todo List</h2>
        <div class="input-group" style="margin-bottom:10px;">
            <input type="text" class="form-control" 
                    placeholder="Please input todos"
                    v-model="name"
                    v-on:keyup.enter="createTodo(name)">
            <span class="input-group-btn">
                <button class="btn btn-default" type="button"
                        @click="createTodo(name)">ADD</button>
            </span>
        </div>
        <ul class="list-group">
            <li class="list-group-item" v-for="(todo, index) in todos" :key="todo">
                {{todo.name}}
                <div class="btn-group pull-right" style="font-size: 12px; line-height: 1;">
                    <button type="button" class="btn-link dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                        More <span class="caret"></span>
                    </button>
                    <ul class="dropdown-menu">
                        <li><a href="#" @click="deleteTodo(todo)">DELETE</a></li> <!--index > todo ?-->
                    </ul>
                </div>
            </li>
        </ul>
    </div>
</template>

<script>
    export default {
        methods:{
            getTodos() {
                var vm = this;
                this.$http.get('http://127.0.0.1:8000/api/todos')
                .then((result) => {
                    console.log(result)
                    vm.todos = result.data.data;
                })
            },
            deleteTodo(i) {
                var vm = this
                this.todos.array.forEach(function(_todo, i, obj) {
                    if(_todo.id === todo.id) {
                        vm.$http.delete('http://127.0.0.1:8000/api/todos/'+todo.id)
                        .then((result)=>{
                            obj.splice(i,1)
                        })
                    }                    
                });
            },
            createTodo(name) {
                if(name != null) {
                    var vm = this;
                    this.$http.defaults.headers.post['Content-Type']='application/json'
                    this.$http.post('http://127.0.0.1:8000/api/todos', {
                        name:name}).then((result) => {
                            vm.todos.push(result.data)
                        })
                    //this.todos.push({name:name});
                    this.name = null
                }
            }
        },
        mounted() {
            this.getTodos();
        },
        name: 'TodoPage',
        data() {
            return {
                todos: [
                    {
                        name:'청소'
                    },
                    {
                        name:'블로그 쓰기'
                    },
                    {
                        name:'밥먹기'
                    },
                    {
                        name:'안녕'
                    }
                ]
            };
        },
        mounted() {
            console.log('Component mounted.')
        }
    }
</script>
