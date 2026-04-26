import {RequireAuth} from "@/components/auth/RequireAuth";

export function DashboardHomePage() {
    return (
        <RequireAuth>
            <p>Dashboard Home</p>
        </RequireAuth>
    )
}