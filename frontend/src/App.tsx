import { Button } from "@/components/ui/button"
import { useEffect, useState } from "react"
import { ReadData, CreateEntry,  DeleteEntry, UpdateEntry } from "../wailsjs/go/main/App"
import { main } from "../wailsjs/go/models"
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { cn } from "./lib/utils"
import { Input } from "./components/ui/input"
import { useToast } from "./hooks/use-toast"
import { Trash } from "lucide-react"
import { Switch } from "./components/ui/switch"

function App() {
  const { toast } = useToast()
  const [todos, setTodos] = useState<main.Todo[]>()
  const [newTodo, setNewTodo] = useState<string>()

  const getTodos = async () => {
    const results = await ReadData()
    setTodos(results)
  }

  const addTodo = async ()=>{
    if(newTodo){
      const res = await CreateEntry(newTodo)
      if(res){
        toast({
          title:"Successfully created"
        })
        setNewTodo(" ")
        await getTodos()
      }else{
        toast({
          title:"Error occured",
          variant:"destructive"
        })
      }
    }
  }

  const updateTodo = async (todo:main.Todo)=>{
    await UpdateEntry(todo.id, todo.name, todo.completed ? 0 : 1)
    await getTodos()
  }

  const deleteTodo = async(id:number)=>{
    await DeleteEntry(id);
    await getTodos()
  }
  useEffect(() => {
    getTodos()
  }, [])

  return (
    <div className="min-h-screen bg-white grid place-items-center mx-auto py-8">
      <div className="text-blue-900 text-2xl font-bold flex flex-col items-center space-y-4">
        <h1>Create new Todo</h1>
        <Input value={newTodo} onChange={(e)=>{setNewTodo(e.target.value)}}/>
        <Button onClick={async () => {await addTodo()}}>Add</Button>
      </div>
      <Table>
        <TableCaption>A list of things you need to do.</TableCaption>
        <TableHeader>
          <TableRow >
            <TableHead>Sl. No.</TableHead>
            <TableHead>Task</TableHead>
            <TableHead>Status</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
            {todos?.map((todo)=>{
              return <TableRow  className={cn(todo.completed ? "line-through bg-green-700" :"bg-red-700", "cursor-pointer")}>
                  <TableCell>{todo.id}</TableCell>
                  <TableCell>{todo.name}</TableCell>
                  <TableCell><Switch onClick={()=>updateTodo(todo)} checked={Boolean(todo.completed)}/></TableCell>
                  <TableCell onClick={async ()=>{await deleteTodo(todo.id)}}><Trash/></TableCell>
                </TableRow>
            })}
        </TableBody>
      </Table>
    </div>
  )
}

export default App
