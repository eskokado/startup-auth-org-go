"use client"
import { useEffect, useState } from "react"
import { Dropdown } from "primereact/dropdown"
import { Card } from "primereact/card"
import { Button } from "primereact/button"
import { billingApi } from "../../services/billing"
import { orgApi, Organization } from "../../services/org"

const plans = [ { label: 'Personal', value: 'PERSONAL' }, { label: 'Organization', value: 'ORGANIZATION' } ]
const cycles = [ { label: 'Mensal', value: 'MONTHLY' }, { label: 'Semestral', value: 'SEMIANNUAL' }, { label: 'Anual', value: 'ANNUAL' } ]

export default function BillingPage() {
  const [plan, setPlan] = useState('PERSONAL')
  const [org, setOrg] = useState<Organization | null>(null)
  const userId = typeof window !== 'undefined' ? localStorage.getItem('user-id') || '' : ''

  useEffect(() => {
    const load = async () => { if (!userId) return; const o = await orgApi.getPersonal(userId); setOrg(o) }
    load()
  }, [userId])

  const startCheckout = async (cycle: string) => {
    if (!org) return
    const res = await billingApi.checkout(org.id, plan, cycle, window.location.origin + '/pages/success', window.location.origin + '/pages/failure')
    window.location.href = res.checkout_url
  }

  return (
    <div className="grid p-4">
      <div className="col-12">
        <Dropdown value={plan} options={plans} onChange={(e) => setPlan(e.value)} placeholder="Plano" />
      </div>
      <div className="col-12 md:col-4">
        <Card title="Mensal">
          <p>Ideal para começar</p>
          <Button label="Assinar" onClick={() => startCheckout('MONTHLY')} />
        </Card>
      </div>
      <div className="col-12 md:col-4">
        <Card title="Semestral">
          <p>Economia no médio prazo</p>
          <Button label="Assinar" onClick={() => startCheckout('SEMIANNUAL')} />
        </Card>
      </div>
      <div className="col-12 md:col-4">
        <Card title="Anual">
          <p>Melhor custo-benefício</p>
          <Button label="Assinar" onClick={() => startCheckout('ANNUAL')} />
        </Card>
      </div>
    </div>
  )
}