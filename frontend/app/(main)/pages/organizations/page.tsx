"use client"
import { useEffect, useState } from "react"
import { InputText } from "primereact/inputtext"
import { Button } from "primereact/button"
import { Card } from "primereact/card"
import { orgApi, Organization } from "../../../services/org"

export default function OrganizationsPage() {
  const [org, setOrg] = useState<Organization | null>(null)
  const [email, setEmail] = useState("")
  const userId = typeof window !== 'undefined' ? localStorage.getItem('user-id') || '' : ''

  useEffect(() => {
    const load = async () => {
      if (!userId) return
      const data = await orgApi.getPersonal(userId)
      setOrg(data)
    }
    load()
  }, [userId])

  const onInvite = async () => {
    if (!org) return
    await orgApi.invite(org.id, userId, email)
    setEmail("")
    alert("Convite enviado")
  }

  return (
    <div className="grid p-4">
      <div className="col-12 md:col-6">
        <Card title="Organização">
          <p>Nome: {org?.name || '...'}</p>
        </Card>
      </div>
      <div className="col-12 md:col-6">
        <Card title="Convidar por e-mail">
          <div className="p-inputgroup">
            <InputText placeholder="email@dominio.com" value={email} onChange={(e) => setEmail(e.target.value)} />
            <Button label="Convidar" onClick={onInvite} disabled={!email || !org} />
          </div>
        </Card>
      </div>
    </div>
  )
}