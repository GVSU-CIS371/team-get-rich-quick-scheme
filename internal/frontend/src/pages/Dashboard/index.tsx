import {DashboardLayout} from "@/components/ui/dashboard-layout";

export function DashboardHomePage() {
    return (
        <DashboardLayout>
            <div className="rounded-2xl border border-zinc-800 bg-zinc-950 p-6">
                <h2 className="text-lg font-medium">Organizations</h2>
                <p className="mt-2 text-sm text-zinc-400">
                    <p>You currently have no organizations</p>
                </p>
            </div>
        </DashboardLayout>
    )
}