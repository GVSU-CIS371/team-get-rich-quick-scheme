import {DashboardLayout} from "@/components/ui/dashboard-layout";
import {useLocation, useRoute} from "preact-iso";
import {useAuth} from "@/components/auth/Auth";
import {useEffect, useState} from "preact/hooks";
import {Button} from "@/components/ui/button";
import {Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";

export function ViewOrganizationPage() {
    const {authClient} = useAuth();
    const location = useLocation();
    const {params} = useRoute();
    const [org, setOrg] = useState(null);
    const [invoices, setInvoices] = useState([]);

    useEffect(() => {
        authClient.get(`/api/v1/organizations/${params.id}`).then(r => {
            if (!r.data.success) {
                location.route('/dashboard');
                return;
            }

            setOrg(r.data.data.organization);

            authClient.get(`/api/v1/organizations/${params.id}/invoices`).then(r => {
                setInvoices(r.data.data.invoices);
            });
        });
    }, []);

    if (!org) {
        return (
            <DashboardLayout>
                <p>Loading...</p>
            </DashboardLayout>
        )
    }

    return (
        <DashboardLayout>
            <div className="rounded-2xl border border-zinc-800 bg-zinc-950 p-6">
                <h2 className="text-lg font-medium">{org.Name}</h2>

                {invoices.length === 0 && (
                    <p className="mt-2 text-sm text-zinc-400">
                        You currently have no created invoices
                    </p>
                )}

                {invoices.length !== 0 && (
                    <div className="mt-4">
                        <Table>
                            <TableCaption>A list of your invoices</TableCaption>
                            <TableHeader>
                                <TableRow>
                                    <TableHead>Date</TableHead>
                                    <TableHead>Note</TableHead>
                                </TableRow>
                            </TableHeader>
                            <TableBody>
                                {invoices.map(inv => (
                                    <TableRow key={inv.ID}>
                                        <TableCell>{(new Date(inv.CreatedAt).toLocaleDateString())}</TableCell>
                                        <TableCell>{inv.Note}</TableCell>
                                        <TableCell>
                                            <Button onClick={() => location.route(
                                                `/dashboard/organizations/${org.ID}/invoices/${inv.ID}`)}>View</Button>
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </div>
                )}

                <Button className="mt-4" onClick={() =>
                    location.route(`/dashboard/organizations/${org.ID}/invoices/add`)}>Create Invoice</Button>
            </div>
        </DashboardLayout>
    )
}