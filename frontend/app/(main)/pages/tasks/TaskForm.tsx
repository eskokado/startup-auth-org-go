"use client"
import { useState } from 'react'
import { Card } from 'primereact/card'
import { InputText } from 'primereact/inputtext'
import { Dropdown } from 'primereact/dropdown'
import { Button } from 'primereact/button'

const statuses = [
  { label: 'TODO', value: 'TODO' },
  { label: 'IN_PROGRESS', value: 'IN_PROGRESS' },
  { label: 'DONE', value: 'DONE' },
]

export default function TaskForm({ title, initial, onSubmit }: { title: string; initial?: { title: string; description: string; status: string }; onSubmit: (v: { title: string; description: string; status: string }) => Promise<void> }) {
  const [t, setT] = useState(initial?.title || '')
  const [d, setD] = useState(initial?.description || '')
  const [s, setS] = useState(initial?.status || 'TODO')
  const [loading, setLoading] = useState(false)

  const submit = async () => { setLoading(true); await onSubmit({ title: t, description: d, status: s }); setLoading(false) }

  return (
    <div className="grid p-4">
      <div className="col-12">
        <Card title={title}>
          <div className="flex gap-2">
            <InputText placeholder="Título" className="w-20rem" value={t} onChange={(e) => setT(e.target.value)} />
            <InputText placeholder="Descrição" className="w-30rem" value={d} onChange={(e) => setD(e.target.value)} />
            <Dropdown value={s} options={statuses} onChange={(e) => setS(e.value)} />
            <Button label="Salvar" onClick={submit} disabled={!t} loading={loading} />
          </div>
        </Card>
      </div>
    </div>
  )
}