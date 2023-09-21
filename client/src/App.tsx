import { ClerkProvider, useAuth } from "@clerk/clerk-react";
function App() {
  const clerkPublishableKey: string | undefined = import.meta.env.VITE_APP_CLERK_PUBLISHABLE_KEY;
  const { userId } = useAuth()
  console.log(clerkPublishableKey)
  return (
	<ClerkProvider publishableKey={clerkPublishableKey as string}>
		<h1>Hi {userId ? userId : ""}</h1>
    
	</ClerkProvider>
  )
}

export default App
