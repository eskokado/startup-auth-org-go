"use client"
import { useEffect, useState } from "react"
import { useRouter } from "next/navigation"
import { InputText } from "primereact/inputtext"
import { Button } from "primereact/button"
import { Dropdown } from "primereact/dropdown"
import { DataTable } from "primereact/datatable"
import { Column } from "primereact/column"
import { Card } from "primereact/card"
import { ConfirmDialog, confirmDialog } from "primereact/confirmdialog"
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
  const userId = typeof window !== 'undefined' ? localStorage.getItem('user-id') || '' : ''
  const router = useRouter()

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

  const statusBody = (row: Task) => (
    <Dropdown value={row.status} options={statuses} onChange={async (e) => { await taskApi.updateStatus(row.id, org!.id, row.title, row.description, e.value); await reload() }} />
  )

  const deleteBody = (row: Task) => (
    <div className="flex gap-2">
      <Button label="Editar" onClick={() => { router.push(`/pages/tasks/${row.id}/edit`) }} />
      <Button label="Excluir" severity="danger" onClick={() => {
        confirmDialog({
          message: `Excluir a tarefa "${row.title}"?`,
          header: "Confirmar exclusão",
          icon: "pi pi-exclamation-triangle",
          acceptClassName: "p-button-danger",
          acceptLabel: "Excluir",
          rejectLabel: "Cancelar",
          accept: async () => { await taskApi.delete(row.id); await reload() },
        })
      }} />
    </div>
  )

  return (
    <div className="grid p-4">
      <div className="col-12 flex justify-content-end">
        <Button label="Adicionar" onClick={() => { router.push('/pages/tasks/new') }} />
      </div>
      <div className="col-12">
        <ConfirmDialog />
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