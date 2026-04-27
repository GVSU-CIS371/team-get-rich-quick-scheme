import {DashboardLayout} from "@/components/ui/dashboard-layout";
import {useLocation, useRoute} from "preact-iso";
import {Button} from "@/components/ui/button";
import {Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {useAuth} from "@/components/auth/Auth";
import {useEffect, useState} from "preact/hooks";

export function ViewInvoicePage() {
    const {authClient} = useAuth();
    const location = useLocation();
    const {params} = useRoute();
    const [invoice, setInvoice] = useState(null);

    useEffect(() => {
        authClient.get(`/api/v1/organizations/${params.id}/invoices/${params.invId}`).then(r => {
            if (!r.data.success) {
                location.route('/dashboard');
                return;
            }

            setInvoice(r.data.data.invoice);
        });
    }, []);

    const deleteItem = (id: number) => {
        authClient.delete(`/api/v1/organizations/${params.id}/invoices/${params.invId}/items/${id}`).then(() => {
            setInvoice(prev => ({
                ...prev,
                Items: prev.Items.filter(i => i.ID !== id)
            }));
        })
    };

    if (!invoice) {
        return (
            <DashboardLayout>
                <p>Loading...</p>
            </DashboardLayout>
        )
    }

    return (
        <DashboardLayout>
            <div className="rounded-2xl border border-zinc-800 bg-zinc-950 p-6">
                <h2 className="text-lg font-medium">Your Invoice</h2>

                {invoice.Items.length === 0 && (
                    <p className="mt-2 text-sm text-zinc-400">
                        You currently have no created invoice items
                    </p>
                )}

                {invoice.Items.length !== 0 && (
                    <div className="mt-4">
                        <Table>
                            <TableCaption>A list of your invoice items</TableCaption>
                            <TableHeader>
                                <TableRow>
                                    <TableHead>Description</TableHead>
                                    <TableHead>Quantity</TableHead>
                                    <TableHead>Price</TableHead>
                                    <TableHead>Total</TableHead>
                                    <TableHead>Actions</TableHead>
                                </TableRow>
                            </TableHeader>
                            <TableBody>
                                {invoice.Items.map(item => (
                                    <TableRow key={item.ID}>
                                        <TableCell>{item.Description}</TableCell>
                                        <TableCell>{item.Quantity}</TableCell>
                                        <TableCell>{item.Price}</TableCell>
                                        <TableCell>{item.Quantity * item.Price}</TableCell>
                                        <TableCell>
                                            <Button onClick={() => deleteItem(item.ID)}>Delete</Button>
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </div>
                )}

                <Button className="mt-4"
                        onClick={() => location.route(
                            `/dashboard/organizations/${params.id}/invoices/${params.invId}/items/add`)}>
                    Create Item
                </Button>

                <Button className="mt-4" onClick={() =>
                    location.route(`/dashboard/organizations/${params.id}`)}>Go Back</Button>
            </div>
        </DashboardLayout>
    )
}