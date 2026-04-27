import {Button} from "@/components/ui/button";
import {RequireAuth} from "@/components/auth/RequireAuth";
import {useLocation} from "preact-iso";

type Link = {
    Name: string;
    Url: string;
}

const links: Link[] = [
    {
        Name: "Organizations",
        Url: "/dashboard"
    },
    {
        Name: "Add Organization",
        Url: "/dashboard/organizations/add"
    },
    {
        Name: "Settings",
        Url: "/dashboard/settings"
    },
    {
        Name: "Logout",
        Url: "/logout"
    }
]

export function DashboardLayout({children}) {
    const location = useLocation();

    return (
        <RequireAuth>
            <div className="min-h-screen bg-black text-white">
                <header className="border-b border-zinc-800 px-4 py-4 sm:px-6 lg:px-8">
                    <div className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
                        <h1 className="text-xl font-semibold tracking-tight">Dashboard</h1>

                        <nav className="grid grid-cols-2 gap-2 sm:flex sm:items-center">
                            {links.map(link => (
                                <Button variant="outline"
                                        className="border-zinc-700 bg-transparent text-white
                                        hover:bg-zinc-900 hover:text-white" onClick={() =>
                                    location.route(link.Url)}>
                                    {link.Name}
                                </Button>
                            ))}
                        </nav>
                    </div>
                </header>

                <main className="px-4 py-8 sm:px-6 lg:px-8">
                    {children}
                </main>
            </div>
        </RequireAuth>
    )
}