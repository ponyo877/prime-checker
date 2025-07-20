import "@typespec/http";

using Http;
@service(#{ title: "Prime Check Service" })
namespace PrimeCheckService;

model PrimeCheck {
  id: int32;
  number: string;
  created_at: utcDateTime;
  trace_id?: string;
  message_id?: string;
  is_prime?: boolean;
  status?: string;
}

model PrimeCheckInput {
  number: string;
}

model PrimeCheckList {
  items: PrimeCheck[];
}

model Setting {
  record_number_success: boolean;
  prime_check_success: boolean;
  email_send_success: boolean;
  dlq_save_success: boolean;
}

@error
model Error {
  code: int32;
  message: string;
}

@route("/prime-check")
@tag("PrimeChecks")
interface PrimeChecks {
  @get get(@path request_id: int32): PrimeCheck | Error;
  @get list(): PrimeCheckList | Error;
  @post create(@body body: PrimeCheckInput): PrimeCheck | Error;
}

@route("/settings")
@tag("Settings")
interface Settings {
  @get get(): Setting | Error;
  @post create(@body body: Setting): Setting | Error;
}
