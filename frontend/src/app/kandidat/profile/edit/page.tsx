"use client"

import type React from "react"
import { useState, useRef } from "react"
import Image from "next/image"
import { useRouter } from "next/navigation"
import Navbar from "@/components/global/navbar"
import Footer from "@/components/global/footer"
import { Button } from "@/components/ui/button"
import { ArrowLeft } from "lucide-react"

interface ProfileData {
  fullName: string
  title: string
  photoUrl: string
  phoneNumber: string
  education: string
  location: string
  cvUrl: string
  cvFile: File | null
  fieldOfInterest: string
}

export default function EditProfileKandidat() {
  const router = useRouter()
  const fileInputRef = useRef<HTMLInputElement>(null)

  const [profile, setProfile] = useState<ProfileData>({
    fullName: "Aditya Pratama",
    title: "Fresh Graduate Teknologi Informasi Universitas Gadjah Mada",
    photoUrl: "/placeholder.svg?height=120&width=120",
    phoneNumber: "081548797755",
    education: "S1",
    location: "Yogyakarta",
    cvUrl: "cv.pdf",
    cvFile: null,
    fieldOfInterest: "IT",
  })

  const [isSaving, setIsSaving] = useState(false)
  const [photoPreview, setPhotoPreview] = useState<string | null>(null)

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setProfile((prev) => ({ ...prev, [name]: value }))
  }

  const handlePhotoClick = () => {
    fileInputRef.current?.click()
  }

  const handlePhotoChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      const file = e.target.files[0]
      const reader = new FileReader()

      reader.onload = (event) => {
        if (event.target?.result) {
          setPhotoPreview(event.target.result as string)
        }
      }

      reader.readAsDataURL(file)
    }
  }

  const handleCVChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      const file = e.target.files[0]
      setProfile((prev) => ({
        ...prev,
        cvFile: file,
        cvUrl: file.name,
      }))
    }
  }

  return (
    <div className="min-h-screen flex flex-col bg-white">
      {/* Header */}
      <Navbar role="kandidat"/>

      {/* Section edit profile */}
      <div className="container mx-auto px-4 py-6">
				{/* Back button */}
        <div className="mb-4">
          <button onClick={() => router.back()} className="flex items-center text-[#1e3a5f] hover:underline">
            <ArrowLeft className="h-4 w-4 mr-1" />
            Kembali
          </button>
        </div>

        <h1 className="text-2xl font-bold text-center text-[#0B2447] mb-6">Ubah Profil</h1>

        <div className="max-w-2xl mx-auto mb-10">
          <form className="bg-[#f8f0e5] rounded-lg border border-[#e8d5c4] p-6 mb-6">
            <div className="flex flex-col items-center mb-6">
              <div className="relative h-32 w-32 rounded-full overflow-hidden mb-6">
                <Image
                  src={photoPreview || profile.photoUrl || "/ProfilePicture.svg"}
                  alt={profile.fullName || "Foto Profil"}
                  fill
                  className="object-cover border border-black"
                />
              </div>
              <button
                type="button"
                onClick={handlePhotoClick}
                className="text-sm text-[#0B2447] border border-[#0B2447] px-3 py-1 rounded-md hover:bg-[#0B2447]/10 cursor-pointer"
              >
                Ubah Foto Profil
              </button>
              <input ref={fileInputRef} type="file" accept="image/*" onChange={handlePhotoChange} className="hidden" />
            </div>

            <div className="space-y-4">
              <div>
                <label className="block font-medium text-gray-700 mb-1">Nama Lengkap</label>
                <input
                  type="text"
                  name="fullName"
                  value={profile.fullName}
                  onChange={handleChange}
                  required
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring focus:ring-[#0B2447] bg-white"
                />
              </div>

              <div>
                <label className="block font-medium text-gray-700 mb-1">Title</label>
                <input
                  type="text"
                  name="title"
                  value={profile.title}
                  onChange={handleChange}
                  required
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring focus:ring-[#0B2447] bg-white"
                />
              </div>

              <div>
                <label className="block font-medium text-gray-700 mb-1">Nomor Telepon</label>
                <input
                  type="tel"
                  name="phoneNumber"
                  value={profile.phoneNumber}
                  onChange={handleChange}
                  required
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring focus:ring-[#0B2447] bg-white"
                />
              </div>

              <div>
                <label className="block font-medium text-gray-700 mb-1">Pendidikan Terakhir</label>
                <input
                  type="text"
                  name="education"
                  value={profile.education}
                  onChange={handleChange}
                  required
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring focus:ring-[#0B2447] bg-white"
                />
              </div>

              <div>
                <label className="block font-medium text-gray-700 mb-1">Bidang yang Diminati</label>
                <input
                  type="text"
                  name="fieldOfInterest"
                  value={profile.fieldOfInterest}
                  onChange={handleChange}
                  required
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring focus:ring-[#0B2447] bg-white"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Lokasi</label>
                <input
                  type="text"
                  name="location"
                  value={profile.location}
                  onChange={handleChange}
                  required
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring focus:ring-[#0B2447] bg-white"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Unggah CV</label>
                <input
                  type="file"
                  accept=".pdf,.doc,.docx"
                  onChange={handleCVChange}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring focus:ring-[#0B2447] bg-white"
                />
                {profile.cvFile && (
                  <p className="mt-2 text-sm text-gray-600">File yang diunggah: {profile.cvUrl}</p>
                )}
              </div>
            </div>
          </form>

          {/* Button to save the edited profile */}
          <div className="flex justify-center mt-8">
            <Button 
              type="submit"
              disabled={isSaving}
              className="bg-[#1e3a5f] text-white px-6 py-2 rounded-md hover:bg-[#1e3a5f]/90 transition-colors disabled:opacity-50 cursor-pointer"
            >
              {isSaving ? "Menyimpan..." : "Simpan profil"}
            </Button>
          </div>
        </div>
      </div>

      {/* Footer */}
      <Footer />
    </div>
  )
}
