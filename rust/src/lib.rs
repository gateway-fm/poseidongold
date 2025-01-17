use plonky2::hash::poseidon_goldilocks::PoseidonGoldHasher;
use std::ffi::c_ulong;

#[no_mangle]
pub extern "C" fn hash12(input: *const c_ulong, capacity: *const c_ulong, output: *mut c_ulong) {
    use plonky2::field::goldilocks_field::GoldilocksField as F;
    use plonky2::field::types::{Field, PrimeField64};
    use plonky2::hash::poseidon::PoseidonHash;

    let mut vals : [u64; 12] = [0; 12];
    unsafe {
        for i in 0..8 {
            vals[i] = *input.add(i) as u64;
        }
        for i in 0..4 {
            vals[i+8] = *capacity.add(i) as u64;
        }
    }
    let mut gl64input = [F::ZERO; 12];
    for i in 0..12 {
        gl64input[i] = F::from_canonical_u64(vals[i]);
    }
    let h = PoseidonHash::hash12(&gl64input);
    unsafe {
        for i in 0..4 {
            *output.add(i) = h[i].to_canonical_u64();
        }
    }
}

#[cfg(test)]
mod tests {
    use std::{ffi::c_ulong, time::Instant};

    use crate::hash12;

    #[test]
    fn test_bench_hash12() {
        let inputs: [c_ulong; 8] = [
            5577006791947779410,
            8674665223082153551,
            15352856648520921629,
            13260572831089785859,
            3916589616287113937,
            6334824724549167320,
            9828766684487745566,
            10667007354186551956,
        ];
        let capacity: [c_ulong; 4] = [0, 0, 0, 0];

        const REPS: usize = 1 << 20;

        let mut output: [c_ulong; 4] = [0, 0, 0, 0];


        let t0 = Instant::now();
        for _i in 0..REPS {
            hash12(inputs.as_ptr(), capacity.as_ptr(), output.as_mut_ptr());
        }
        let tdiff = Instant::now().duration_since(t0).as_nanos();

        assert_eq!(output[0], 7986352640330579808);
        assert_eq!(output[1], 16698976638447200418);
        assert_eq!(output[2], 14099060853601989680);
        assert_eq!(output[3], 1806029100513259151);

        println!("Time in s is: {:?}", tdiff as f64 / (1000000000 as f64));
        println!("Time per hash in ns is: {:?}", tdiff as f64 / (REPS as f64));
    }
}
