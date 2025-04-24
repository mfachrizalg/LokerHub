// src/hooks/tambahLowongan.ts
import { useState } from "react"
import { useRouter } from "next/navigation"

export interface JobFormData {
  title: string
  jobType: string
  positionLevel: string
  salaryMin: string
  salaryMax: string
  field: string
  closingDate: string
  description: string
  responsibilities: string[]
  qualifications: string[]
}

export function addJobForm(initialFormData: any) {
  const router = useRouter()
  const [formData, setFormData] = useState<JobFormData>({
    title: "",
    jobType: "",
    positionLevel: "",
    salaryMin: "",
    salaryMax: "",
    field: "",
    closingDate: "",
    description: "",
    responsibilities: [""],
    qualifications: [""],
})
  const [newResponsibility, setNewResponsibility] = useState("")
  const [newQualification, setNewQualification] = useState("")
  const [isSubmitting, setIsSubmitting] = useState(false)

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target
    setFormData((prev: JobFormData) => ({ ...prev, [name]: value }))
  }

  const addResponsibility = () => {
    if (newResponsibility.trim() !== "") {
      setFormData((prev: JobFormData) => ({
        ...prev,
        responsibilities: [...prev.responsibilities, newResponsibility],
      }))
      setNewResponsibility("")
    }
  }

  const addQualification = () => {
    if (newQualification.trim() !== "") {
      setFormData((prev: JobFormData) => ({
        ...prev,
        qualifications: [...prev.qualifications, newQualification],
      }))
      setNewQualification("")
    }
  }

  const removeResponsibility = (index: number) => {
    setFormData((prev: JobFormData) => ({
      ...prev,
      responsibilities: prev.responsibilities.filter((_, i) => i !== index),
    }))
  }

  const removeQualification = (index: number) => {
    setFormData((prev: JobFormData) => ({
      ...prev,
      qualifications: prev.qualifications.filter((_, i) => i !== index),
    }))
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsSubmitting(true)

    try {
      const cleanedData = {
        ...formData,
        responsibilities: formData.responsibilities.filter((item: string) => item.trim() !== ""),
        qualifications: formData.qualifications.filter((item: string) => item.trim() !== ""),
      }

      console.log("Submitting job:", cleanedData)

      await new Promise((resolve) => setTimeout(resolve, 1000))
      alert("Lowongan berhasil ditambahkan!")
      router.push("/")
    } catch (error) {
      console.error("Error adding job:", error)
      alert("Gagal menambahkan lowongan. Silakan coba lagi.")
    } finally {
      setIsSubmitting(false)
    }
  }

  return {
    formData,
    setFormData,
    newResponsibility,
    setNewResponsibility,
    newQualification,
    setNewQualification,
    isSubmitting,
    handleChange,
    addResponsibility,
    addQualification,
    removeResponsibility,
    removeQualification,
    handleSubmit,
  }
}
