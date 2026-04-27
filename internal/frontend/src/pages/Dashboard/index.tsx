import {DashboardLayout} from "@/components/ui/dashboard-layout";
import {useAuth} from "@/components/auth/Auth";
import {useState} from "preact/hooks";
import {Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {Button} from "@/components/ui/button";
import {useLocation} from "preact-iso";

export function DashboardHomePage() {
    const {authClient} = useAuth();
    const [organizations, setOrganizations] = useState([]);
    const location = useLocation();

    authClient.get('/api/v1/organizations').then(r => {
       setOrganizations(r.data.data.organizations);
    });

    return (
        <DashboardLayout>
            <div className="rounded-2xl border border-zinc-800 bg-zinc-950 p-6">
                <h2 className="text-lg font-medium">Organizations</h2>
                {organizations.length === 0 && (
                    <p className="mt-2 text-sm text-zinc-400">
                        You currently have no organizations
                    </p>
                )}

                {organizations.length !== 0 && (
                    <div className="mt-4">
                        <Table>
                            <TableCaption>A list of your organizations</TableCaption>
                            <TableHeader>
                                <TableRow>
                                    <TableHead>Name</TableHead>
                                    <TableHead>Description</TableHead>
                                    <TableHead>Actions</TableHead>
                                </TableRow>
                            </TableHeader>
                            <TableBody>
                                {organizations.map(org => (
                                   <TableRow key={org.ID}>
                                       <TableCell>{org.Name}</TableCell>
                                       <TableCell>{org.Description}</TableCell>
                                       <TableCell>
                                           <Button onClick={() =>
                                               location.route(`/dashboard/organizations/${org.ID}`)}>View</Button>
                                       </TableCell>
                                   </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </div>
                )}
            </div>
        </DashboardLayout>
    )
}