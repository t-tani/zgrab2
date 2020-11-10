# zschema sub-schema for zgrab2's fox module
# Registers zgrab2-quasar globally, and quasar with the main zgrab2 schema.
from zschema.leaves import *
from zschema.compounds import *
import zschema.registry

import zcrypto_schemas.zcrypto as zcrypto
from . import zgrab2

quasar_scan_response = SubRecord({
    'result': SubRecord({
        'is_quasar': Boolean(),
        'version': String(),
        # 'id': Unsigned32BitInteger(),
        # 'hostname': String(),
        # 'host_address': String(),
    })
}, extends=zgrab2.base_scan_response)

zschema.registry.register_schema('zgrab2-quasar', quasar_scan_response)

zgrab2.register_scan_response_type('quasar', quasar_scan_response)
