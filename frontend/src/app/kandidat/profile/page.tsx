"use client"

import Image from "next/image"
import { useRouter } from "next/navigation"

const profile = {
  fullName: "Aditya Pratama",
  title: "Fresh Graduate Teknologi Informasi Universitas Gadjah Mada",
  photoUrl: "/placeholder.svg?height=120&width=120",
  phoneNumber: "081548797755",
  education: "S1",
  location: "Yogyakarta",
  cvUrl: "cv.pdf",
  fieldOfInterest: "IT"
}

export default function ProfileKandidat() {
  const router = useRouter()

  const handleEditProfile = () => {
    router.push('profile/edit')
  }

  return (
    <div className="min-h-screen flex flex-col bg-white">
      {/* Header */}
      <header className="border-b">
        <div className="container mx-auto py-6 flex items-center justify-between">
          <div className="flex items-center">
            <Image
              src="/LokerHub_Logo.svg"
              alt="logo"
              width={222}
              height={42}
              className="w-full "
              />
          </div>
        </div>
      </header>

      {/* Profile section */}
      <div className="container mx-auto px-4 py-8">
        <h1 className="text-4xl font-bold text-center text-[#1e3a5f] mb-8">Profil</h1>
        <div className="max-w-2xl mx-auto">
          <div className="bg-[#f5efe7]/40 rounded-lg border border-blue-200 py-6">
            <div className="flex flex-col items-center mb-6">
              <div className="relative h-32 w-32 rounded-full border border-black overflow-hidden mb-4">
                <Image 
                  src="/ProfilePicture.svg" 
                  alt={profile.fullName}
                  fill
                  className="object-cover"
                />
              </div>
              <h2 className="text-xl font-bold text-[#1e3a5f]">{profile.fullName}</h2>
              <p className="text-lg text-center text-gray-600">{profile.title}</p>
            </div>
            
            <div className="grid grid-cols-2 gap-4 mb-2 text-center">
              <div>
                <p className="text-gray-500">Nomor Telepon</p>
                <p className="font-medium">{profile.phoneNumber}</p>
              </div>
              <div>
                <p className="text-gray-500">Pendidikan Terakhir</p>
                <p className="font-medium">{profile.education}</p>
              </div>
              <div>
                <p className="text-gray-500">Lokasi</p>
                <p className="font-medium">{profile.location}</p>
              </div>
              <div>
                <p className="text-gray-500">CV</p>
                <a 
                  href={profile.cvUrl} 
                  target="_blank" 
                  rel="noopener noreferrer"
                  className="inline-block border border-black rounded-lg text-gray-700 px-3 py-1 hover:bg-blue-200"
                >
                  cv.pdf
                </a>
              </div>
            </div>  
          </div>

					<div className="flex justify-center pt-10">
            <button
                onClick={handleEditProfile}
                className="border border-[#1e3a5f] bg-[#f5efe7]/40 text-[#1e3a5f] font-semibold px-4 py-1 rounded-md hover:bg-[#1e3a5f]/10 transition-colors cursor-pointer"
              >
                Ubah profil
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}