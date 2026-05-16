import { useEffect } from "react"
import { useNavigate } from "react-router-dom"

export function LoginCallbackPage() {
    const navigate = useNavigate()

    useEffect(() => {
        navigate("/")
    }, [navigate])

    return <></>
}