import { render } from 'preact';
import { LocationProvider, Router, Route } from 'preact-iso';

import { HomePage } from './pages/Home';
import { NotFound } from './pages/_404.jsx';
import './style.css';
import {LoginPage} from "@/pages/Login";
import {RegisterPage} from "./pages/Register";
import {AuthProvider} from "@/components/auth/Auth";
import {DashboardHomePage} from "@/pages/Dashboard";
import {LogoutPage} from "@/pages/Logout";
import {AddOrganizationPage} from "@/pages/Dashboard/organizations/add";
import {ViewOrganizationPage} from "@/pages/Dashboard/organizations/id";
import {AddInvoicePage} from "@/pages/Dashboard/organizations/invoices/add";
import {ViewInvoicePage} from "@/pages/Dashboard/organizations/invoices/id";
import {AddInvoiceItemPage} from "@/pages/Dashboard/organizations/invoices/id/item/add";

export function App() {
	return (
		<AuthProvider>
			<div className="dark bg-background min-h-screen">
				<LocationProvider>
					<main>
						<Router>
							<Route path="/" component={HomePage} />
							<Route path="/login" component={LoginPage} />
							<Route path="/register" component={RegisterPage} />
							<Route path="/logout" component={LogoutPage} />
							<Route path="/dashboard" component={DashboardHomePage} />
							<Route path="/dashboard/organizations/add" component={AddOrganizationPage} />
							<Route path="/dashboard/organizations/:id" component={ViewOrganizationPage} />
							<Route path="/dashboard/organizations/:id/invoices/add" component={AddInvoicePage} />
							<Route path="/dashboard/organizations/:id/invoices/:invId" component={ViewInvoicePage} />
							<Route path="/dashboard/organizations/:id/invoices/:invId/items/add"
								   component={AddInvoiceItemPage} />
							<Route default component={NotFound} />
						</Router>
					</main>
				</LocationProvider>
			</div>
		</AuthProvider>
	);
}

render(<App />, document.getElementById('app'));
