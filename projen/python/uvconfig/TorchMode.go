package uvconfig


// The strategy to use when determining the appropriate PyTorch index.
// Experimental.
type TorchMode string

const (
	// Select the appropriate PyTorch index based on the operating system and CUDA driver version.
	//
	// (auto).
	// Experimental.
	TorchMode_AUTO TorchMode = "AUTO"
	// Use the CPU-only PyTorch index.
	//
	// (cpu).
	// Experimental.
	TorchMode_CPU TorchMode = "CPU"
	// Use the PyTorch index for CUDA 13.0. (cu130).
	// Experimental.
	TorchMode_CU130 TorchMode = "CU130"
	// Use the PyTorch index for CUDA 12.9. (cu129).
	// Experimental.
	TorchMode_CU129 TorchMode = "CU129"
	// Use the PyTorch index for CUDA 12.8. (cu128).
	// Experimental.
	TorchMode_CU128 TorchMode = "CU128"
	// Use the PyTorch index for CUDA 12.6. (cu126).
	// Experimental.
	TorchMode_CU126 TorchMode = "CU126"
	// Use the PyTorch index for CUDA 12.5. (cu125).
	// Experimental.
	TorchMode_CU125 TorchMode = "CU125"
	// Use the PyTorch index for CUDA 12.4. (cu124).
	// Experimental.
	TorchMode_CU124 TorchMode = "CU124"
	// Use the PyTorch index for CUDA 12.3. (cu123).
	// Experimental.
	TorchMode_CU123 TorchMode = "CU123"
	// Use the PyTorch index for CUDA 12.2. (cu122).
	// Experimental.
	TorchMode_CU122 TorchMode = "CU122"
	// Use the PyTorch index for CUDA 12.1. (cu121).
	// Experimental.
	TorchMode_CU121 TorchMode = "CU121"
	// Use the PyTorch index for CUDA 12.0. (cu120).
	// Experimental.
	TorchMode_CU120 TorchMode = "CU120"
	// Use the PyTorch index for CUDA 11.8. (cu118).
	// Experimental.
	TorchMode_CU118 TorchMode = "CU118"
	// Use the PyTorch index for CUDA 11.7. (cu117).
	// Experimental.
	TorchMode_CU117 TorchMode = "CU117"
	// Use the PyTorch index for CUDA 11.6. (cu116).
	// Experimental.
	TorchMode_CU116 TorchMode = "CU116"
	// Use the PyTorch index for CUDA 11.5. (cu115).
	// Experimental.
	TorchMode_CU115 TorchMode = "CU115"
	// Use the PyTorch index for CUDA 11.4. (cu114).
	// Experimental.
	TorchMode_CU114 TorchMode = "CU114"
	// Use the PyTorch index for CUDA 11.3. (cu113).
	// Experimental.
	TorchMode_CU113 TorchMode = "CU113"
	// Use the PyTorch index for CUDA 11.2. (cu112).
	// Experimental.
	TorchMode_CU112 TorchMode = "CU112"
	// Use the PyTorch index for CUDA 11.1. (cu111).
	// Experimental.
	TorchMode_CU111 TorchMode = "CU111"
	// Use the PyTorch index for CUDA 11.0. (cu110).
	// Experimental.
	TorchMode_CU110 TorchMode = "CU110"
	// Use the PyTorch index for CUDA 10.2. (cu102).
	// Experimental.
	TorchMode_CU102 TorchMode = "CU102"
	// Use the PyTorch index for CUDA 10.1. (cu101).
	// Experimental.
	TorchMode_CU101 TorchMode = "CU101"
	// Use the PyTorch index for CUDA 10.0. (cu100).
	// Experimental.
	TorchMode_CU100 TorchMode = "CU100"
	// Use the PyTorch index for CUDA 9.2. (cu92).
	// Experimental.
	TorchMode_CU92 TorchMode = "CU92"
	// Use the PyTorch index for CUDA 9.1. (cu91).
	// Experimental.
	TorchMode_CU91 TorchMode = "CU91"
	// Use the PyTorch index for CUDA 9.0. (cu90).
	// Experimental.
	TorchMode_CU90 TorchMode = "CU90"
	// Use the PyTorch index for CUDA 8.0. (cu80).
	// Experimental.
	TorchMode_CU80 TorchMode = "CU80"
	// Use the PyTorch index for ROCm 6.3. (rocm6.3).
	// Experimental.
	TorchMode_ROCM6_3 TorchMode = "ROCM6_3"
	// Use the PyTorch index for ROCm 6.2.4. (rocm6.2.4).
	// Experimental.
	TorchMode_ROCM6_2_4 TorchMode = "ROCM6_2_4"
	// Use the PyTorch index for ROCm 6.2. (rocm6.2).
	// Experimental.
	TorchMode_ROCM6_2 TorchMode = "ROCM6_2"
	// Use the PyTorch index for ROCm 6.1. (rocm6.1).
	// Experimental.
	TorchMode_ROCM6_1 TorchMode = "ROCM6_1"
	// Use the PyTorch index for ROCm 6.0. (rocm6.0).
	// Experimental.
	TorchMode_ROCM6_0 TorchMode = "ROCM6_0"
	// Use the PyTorch index for ROCm 5.7. (rocm5.7).
	// Experimental.
	TorchMode_ROCM5_7 TorchMode = "ROCM5_7"
	// Use the PyTorch index for ROCm 5.6. (rocm5.6).
	// Experimental.
	TorchMode_ROCM5_6 TorchMode = "ROCM5_6"
	// Use the PyTorch index for ROCm 5.5. (rocm5.5).
	// Experimental.
	TorchMode_ROCM5_5 TorchMode = "ROCM5_5"
	// Use the PyTorch index for ROCm 5.4.2. (rocm5.4.2).
	// Experimental.
	TorchMode_ROCM5_4_2 TorchMode = "ROCM5_4_2"
	// Use the PyTorch index for ROCm 5.4. (rocm5.4).
	// Experimental.
	TorchMode_ROCM5_4 TorchMode = "ROCM5_4"
	// Use the PyTorch index for ROCm 5.3. (rocm5.3).
	// Experimental.
	TorchMode_ROCM5_3 TorchMode = "ROCM5_3"
	// Use the PyTorch index for ROCm 5.2. (rocm5.2).
	// Experimental.
	TorchMode_ROCM5_2 TorchMode = "ROCM5_2"
	// Use the PyTorch index for ROCm 5.1.1. (rocm5.1.1).
	// Experimental.
	TorchMode_ROCM5_1_1 TorchMode = "ROCM5_1_1"
	// Use the PyTorch index for ROCm 4.2. (rocm4.2).
	// Experimental.
	TorchMode_ROCM4_2 TorchMode = "ROCM4_2"
	// Use the PyTorch index for ROCm 4.1. (rocm4.1).
	// Experimental.
	TorchMode_ROCM4_1 TorchMode = "ROCM4_1"
	// Use the PyTorch index for ROCm 4.0.1. (rocm4.0.1).
	// Experimental.
	TorchMode_ROCM4_0_1 TorchMode = "ROCM4_0_1"
	// Use the PyTorch index for Intel XPU.
	//
	// (xpu).
	// Experimental.
	TorchMode_XPU TorchMode = "XPU"
)

