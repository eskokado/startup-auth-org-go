"use client"
import { useEffect, useState } from "react"
import { InputText } from "primereact/inputtext"
import { Button } from "primereact/button"
import { Dropdown } from "primereact/dropdown"
import { DataTable } from "primereact/datatable"
import { Column } from "primereact/column"
import { Card } from "primereact/card"
import { taskApi, Task } from "../../../services/tasks"
import { orgApi, Organization } from "../../../services/org"

const statuses = [
  { label: 'TODO', value: 'TODO' },
  { label: 'IN_PROGRESS', value: 'IN_PROGRESS' },
  { label: 'DONE', value: 'DONE' },
]

export default function TasksPage() {
  const [org, setOrg] = useState<Organization | null>(null)
  const [tasks, setTasks] = useState<Task[]>([])
  const [title, setTitle] = useState("")
  const [description, setDescription] = useState("")
  const userId = typeof window !== 'undefined' ? localStorage.getItem('user-id') || '' : ''

  useEffect(() => {
    const load = async () => {
      if (!userId) return
      const data = await orgApi.getPersonal(userId)
      setOrg(data)
      const list = await taskApi.list(data.id)
      setTasks(list)
    }
    load()
  }, [userId])

  const reload = async () => {
    if (!org) return
    const list = await taskApi.list(org.id)
    setTasks(list)
  }

  const onCreate = async () => {
    if (!org || !title) return
    await taskApi.create(org.id, title, description)
    setTitle("")
    setDescription("")
    await reload()
  }

  const statusBody = (row: Task) => (
    <Dropdown value={row.status} options={statuses} onChange={async (e) => {
      await taskApi.updateStatus(row.id, org!.id, row.title, row.description, e.value)
      await reload()
    }} />
  )

  const deleteBody = (row: Task) => (
    <Button label="Excluir" severity="danger" onClick={async () => { await taskApi.delete(row.id); await reload() }} />
  )

  return (
    <div className="grid p-4">
      <div className="col-12">
        <Card title="Nova tarefa">
          <div className="flex gap-2">
            <InputText placeholder="Título" value={title} onChange={(e) => setTitle(e.target.value)} />
            <InputText placeholder="Descrição" value={description} onChange={(e) => setDescription(e.target.value)} />
            <Button label="Criar" onClick={onCreate} disabled={!org || !title} />
          </div>
        </Card>
      </div>
      <div className="col-12">
        <DataTable value={tasks} tableStyle={{ minWidth: '60rem' }}>
          <Column field="title" header="Título" />
          <Column field="description" header="Descrição" />
          <Column header="Status" body={statusBody} />
          <Column header="Ações" body={deleteBody} />
        </DataTable>
      </div>
    </div>
  )
}