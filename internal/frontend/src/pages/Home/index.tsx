import {Button} from "@/components/ui/button";
import {Card, CardContent} from "@/components/ui/card";
import {useAuth} from "@/components/auth/Auth";
import {useState} from "preact/hooks";
import {useLocation} from "preact-iso";

export function HomePage() {
	const {authClient} = useAuth();
	const [activeUsers, setActiveUsers] = useState(0);
	const [invoiceCount, setInvoiceCount] = useState(0);
	const location = useLocation();

	authClient.get("/api/v1/stats").then(r => {
		setActiveUsers(r.data.data.userCount);
		setInvoiceCount(r.data.data.invoiceCount)
	})

	return (
		<div className="min-h-screen text-white">
			<div className="flex justify-between items-center px-8 py-4">
				<h1 className="text-xl font-semibold">Invoice Gen</h1>
				<Button variant="secondary" className="bg-white text-black hover:bg-gray-200"
						onClick={() => location.route('/login')}>
					Login
				</Button>
			</div>

			<div className="flex flex-col items-center justify-center text-center mt-20 px-6">
				<h2 className="text-4xl font-bold mb-4">Simple Invoice Generation</h2>
				<p className="text-gray-400 max-w-xl">
					Create, manage, and send professional invoices in seconds. Built for freelancers and small businesses.
				</p>
			</div>

			<div className="grid grid-cols-1 md:grid-cols-2 gap-6 mt-16 px-8 max-w-4xl mx-auto">
				<Card className="bg-zinc-900 border border-zinc-800 rounded-2xl shadow-lg">
					<CardContent className="p-6">
						<p className="text-gray-400 text-sm">Invoices Generated</p>
						<h3 className="text-3xl font-bold mt-2">{invoiceCount}</h3>
					</CardContent>
				</Card>

				<Card className="bg-zinc-900 border border-zinc-800 rounded-2xl shadow-lg">
					<CardContent className="p-6">
						<p className="text-gray-400 text-sm">Active Users</p>
						<h3 className="text-3xl font-bold mt-2">{activeUsers}</h3>
					</CardContent>
				</Card>
			</div>
		</div>
	);
}

